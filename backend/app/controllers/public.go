package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/middlewares"
	"backend/pkg/utils"
	"backend/platform/database"
)

func GetClasses(c *fiber.Ctx) error {
	tables := []string{"categories", "conclusions", "roles", "groups", "statuses", "regions"}
	models := []interface{}{&[]models.Category{}, &[]models.Conclusion{}, &[]models.Role{}, &[]models.Group{}, &[]models.Status{}, &[]models.Region{}}

	db := database.OpenDb()
	results := make([]map[string]interface{}, len(tables))

	for i, table := range tables {
		db.Table(table).Find(models[i])
		results[i] = map[string]interface{}{table: models[i]}
	}

	return c.Status(200).JSON(results)
}

// indexHandler handles the request to the index route.
func PostIndex(c *fiber.Ctx) error {
	payload := struct {
		Search string `json:"search"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		log.Println(err)
	}

	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}

	db := database.OpenDb()
	var persons []models.Person
	var pagination = 16
	var hasPrev, hasNext bool

	switch c.Params("item") {
	case "new":
		statusNew := models.Status{}.GetID(utils.Statuses["new"])
		statusUpd := models.Status{}.GetID(utils.Statuses["update"])
		statusRep := models.Status{}.GetID(utils.Statuses["repeat"])
		db.
			Where("status_id = ? OR status_id = ? OR status_id = ?", statusNew, statusUpd, statusRep).
			Find(&persons).
			Limit(pagination).
			Offset(pagination * (intPage - 1)).
			Order("created_at desc")
	case "officer":
		var checks []models.Check
		statusFin := models.Status{}.GetID(utils.Statuses["finish"])
		statusCan := models.Status{}.GetID(utils.Statuses["cancel"])
		db.
			Find(&persons).
			Where("status_id != ? OR status_id != ?", statusFin, statusCan).
			Limit(pagination).
			Offset(pagination * (intPage - 1)).
			Order("created_at desc").
			Joins("JOIN checks ON checks.person_id = people.id").
			Where(models.Check{Officer: "current"}).
			Find(&checks)
	case "search":
		db.
			Where("fullname LIKE ?", "%"+payload.Search+"%").
			Find(&persons).Limit(10).
			Offset(10 * (intPage - 1)).
			Order("created_at desc")
	default:
		return nil
	}

	if intPage > 1 {
		hasPrev = true
	}
	if len(persons) == pagination {
		hasNext = true
	}
	result, err := json.Marshal(persons)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{"result": result, "hasNext": hasNext, "hasPrev": hasPrev})
}

func GetResume(c *fiber.Ctx) error {
	db := database.OpenDb()
	var person models.Person

	db.Where("id = ?", c.Params("person_id")).First(&person)

	switch c.Params("action") {
	case "status":
		person.StatusID = models.Status{}.GetID(utils.Statuses["update"])
		db.Save(&person)
		return c.Status(200).JSON(person)

	case "send":
		status := models.Status{}
		if person.StatusID == status.GetID(utils.Statuses["new"]) ||
			person.StatusID == status.GetID(utils.Statuses["update"]) ||
			person.StatusID == status.GetID(utils.Statuses["repeat"]) {

			docs := models.Document{}
			db.
				Where("person_id = ?", person.ID).
				Order("created_at desc").
				First(&docs)

			addr := models.Address{}
			db.
				Where("person_id = ? view LIKE ?", person.ID, "%регистрации%").
				Order("created_at desc").
				First(&addr)

			bodyStruct := struct {
				Person models.Person
				Docs   models.Document
				Addr   models.Address
			}{
				Person: person,
				Docs:   docs,
				Addr:   addr,
			}

			jsonBody, err := json.Marshal(bodyStruct)
			if err != nil {
				return c.Status(500).JSON(err)
			}

			agent := fiber.Post("https://httpbin.org/get")
			agent.Body(jsonBody)
			statusCode, _, errs := agent.Bytes()
			if len(errs) > 0 {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"errs": errs,
				})
			}
			if statusCode != 200 {
				return c.Status(500).JSON(err)
			} else {
				person.StatusID = models.Status{}.GetID(utils.Statuses["finish"])
				db.Save(&person)
				return c.Status(200).JSON(person)
			}
		}
	}
	return c.Status(200).JSON(person)
}

func DeleteResume(c *fiber.Ctx) error {
	db := database.OpenDb()
	var person models.Person

	db.Where("id = ?", c.Params("person_id")).First(&person)
	err := os.RemoveAll(person.PathToDocs)
	if err != nil {
		log.Println(err)
	}
	db.Delete(&person)

	return c.Status(204).JSON("Person deleted")
}

func PostResume(c *fiber.Ctx) error {
	db := database.OpenDb()
	var person models.Person
	var resume models.Person

	err := c.BodyParser(&resume)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Where("fullname = ?, birthday = ?", c.Params("person_id"), c.Params("birthday")).
		First(&person)

	if person.ID == 0 {
		resume.StatusID = models.Status{}.GetID(utils.Statuses["new"])
		db.Create(&resume)
		person.ID = resume.ID
	} else {
		resume.StatusID = models.Status{}.GetID(utils.Statuses["update"])
		db.Save(&resume)
	}
	person.PathToDocs = makeFolder(person.FullName, person.ID)
	db.Save(&person)

	return c.Status(200).JSON(person.ID)
}

func makeFolder(fullname string, person_id uint) string {
	path := filepath.Join(strings.ToUpper(string(fullname[0])), fmt.Sprintf("%d-%s", person_id, fullname))
	basePath := os.Getenv("BASE_PATH")
	url := filepath.Join(basePath, path)
	stat, err := os.Stat(url)
	if err == nil && !stat.IsDir() {
		os.Mkdir(url, 0755)
	}
	return path
}

// Staffs Handlers
func GetStaffs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var staffs []models.Staff
	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&staffs)

	return c.Status(200).JSON(staffs)
}

func PostStaffs(c *fiber.Ctx) error {

	db := database.OpenDb()
	var staff models.Staff

	err := c.BodyParser(&staff)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	staff.PersonID = uint(itemID)
	db.Create(&staff)

	return c.Status(200).JSON("Created")
}

func PatchStaffs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var staff models.Staff

	err := c.BodyParser(&staff)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&staff).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&staff)

	return c.Status(200).JSON("Updated")
}

func DeleteStaffs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var staff models.Staff

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&staff)

	db.Delete(&staff)

	return c.Status(200).JSON("Deleted")
}

// Documents Handlers
func GetDocs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var documents []models.Document

	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&documents)

	return c.Status(200).JSON(documents)
}

func PostDocs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var document models.Document

	err := c.BodyParser(&document)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	document.PersonID = uint(itemID)

	db.Create(&document)

	return c.Status(200).JSON("Created")
}

func PatchDocs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var document models.Document

	err := c.BodyParser(&document)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&document).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&document)

	return c.Status(200).JSON("Updated")
}

func DeleteDocs(c *fiber.Ctx) error {
	db := database.OpenDb()
	var document models.Document

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&document)

	db.Delete(&document)

	return c.Status(200).JSON("Deleted")
}

// Address Handlers
func GetAddress(c *fiber.Ctx) error {

	db := database.OpenDb()
	var address []models.Address

	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&address)

	return c.Status(200).JSON(address)
}

func PostAddress(c *fiber.Ctx) error {
	db := database.OpenDb()
	var address models.Address

	err := c.BodyParser(&address)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	address.PersonID = uint(itemID)

	db.Create(&address)

	return c.Status(200).JSON("Created")
}

func PatchAddress(c *fiber.Ctx) error {
	db := database.OpenDb()
	var address models.Address

	err := c.BodyParser(&address)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&address).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&address)

	return c.Status(200).JSON("Updated")
}

func DeleteAddress(c *fiber.Ctx) error {
	db := database.OpenDb()
	var address models.Address

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&address)

	db.Delete(&address)

	return c.Status(200).JSON("Deleted")
}

// Contact Handlers
func GetContact(c *fiber.Ctx) error {
	db := database.OpenDb()
	var contact []models.Contact

	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&contact)

	return c.Status(200).JSON(contact)
}

func PostContact(c *fiber.Ctx) error {
	db := database.OpenDb()
	var contact models.Contact

	err := c.BodyParser(&contact)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	contact.PersonID = uint(itemID)

	db.Create(&contact)

	return c.Status(200).JSON("Created")
}

func PatchContact(c *fiber.Ctx) error {
	db := database.OpenDb()
	var contact models.Contact

	err := c.BodyParser(&contact)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&contact).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&contact)

	return c.Status(200).JSON("Updated")
}

func DeleteContact(c *fiber.Ctx) error {
	db := database.OpenDb()
	var contact models.Contact

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&contact)

	db.Delete(&contact)

	return c.Status(200).JSON("Deleted")
}

// Work Handlers
func GetWorkplace(c *fiber.Ctx) error {
	db := database.OpenDb()
	var workplace []models.Workplace

	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&workplace)

	return c.Status(200).JSON(workplace)
}

func PostWorkplace(c *fiber.Ctx) error {
	db := database.OpenDb()
	var workplace models.Workplace

	err := c.BodyParser(&workplace)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	workplace.PersonID = uint(itemID)

	db.Create(&workplace)

	return c.Status(200).JSON("Created")
}

func PatchWorkplace(c *fiber.Ctx) error {
	db := database.OpenDb()
	var workplace models.Workplace

	err := c.BodyParser(&workplace)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&workplace).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&workplace)

	return c.Status(200).JSON("Updated")
}

func DeleteWorkplace(c *fiber.Ctx) error {
	db := database.OpenDb()
	var workplace models.Workplace

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&workplace)

	db.Delete(&workplace)

	return c.Status(200).JSON("Deleted")
}

// Affiliations Handlers
func GetAffilation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var affilation []models.Affilation

	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&affilation)

	return c.Status(200).JSON(affilation)
}

func PostAffilation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var affilation models.Affilation

	err := c.BodyParser(&affilation)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	affilation.PersonID = uint(itemID)

	db.Create(&affilation)

	return c.Status(200).JSON("Created")
}

func PatchAffilation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var affilation models.Affilation

	err := c.BodyParser(&affilation)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&affilation).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&affilation)

	return c.Status(200).JSON("Updated")
}

func DeleteAffilation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var affilation models.Affilation

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&affilation)

	db.Delete(&affilation)

	return c.Status(200).JSON("Deleted")
}

// Relationships Handlers
func GetRelation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var relation []models.Relation

	db.
		Where("item_id = ?", c.Params("item_id")).
		Find(&relation)

	return c.Status(200).JSON(relation)
}

func PostRelation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var relation models.Relation

	err := c.BodyParser(&relation)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	relation.PersonID = uint(itemID)
	db.Create(&relation)

	reRelation := models.Relation{
		View:     relation.View,
		Relation: relation.ID,
		PersonID: relation.PersonID,
	}

	db.Create(&reRelation)

	return c.Status(201).JSON("Created")
}

func PatchRelation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var relation models.Relation

	err := c.BodyParser(&relation)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&relation).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&relation)

	return c.Status(200).JSON("Updated")
}

func DeleteRelation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var relation models.Relation

	db.
		Where("item_id = ?", c.Params("item_id")).
		First(&relation)

	db.Delete(&relation)

	return c.Status(200).JSON("Deleted")
}

// Check Handlers
func GetCheck(c *fiber.Ctx) error {
	var check models.Check
	var checks []models.Check
	db := database.OpenDb()
	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	if c.Params("action") == "add" {

		var person models.Person
		db.First(&person, c.Params("item_id"))
		model := models.Status{}
		if person.StatusID == model.GetID(utils.Statuses["new"]) ||
			person.StatusID == model.GetID(utils.Statuses["update"]) ||
			person.StatusID == model.GetID(utils.Statuses["repeat"]) {

			person.StatusID = model.GetID(utils.Statuses["manual"])
			db.Save(&person)

			itemIDStr := c.Params("item_id")
			itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
			if err != nil {
				return c.Status(500).JSON(err)
			}

			check.Officer = tokenMeta.FullName
			check.PersonID = uint(itemID)
			db.Create(&check)
		}
		return c.Status(200).JSON("Added")

	} else if c.Params("action") == "self" {
		db.First(&check, c.Params("item_id"))

		var oldUser models.User
		db.
			Where("fullname = ?", check.Officer).
			First(&oldUser)

		var message models.Message
		if oldUser.ID != tokenMeta.UserID {
			message.MessageContent = "Aнкета переделегирована " + tokenMeta.FullName
			message.UserID = oldUser.ID
			db.Create(&message)
			message.UserID = tokenMeta.UserID
			db.Create(&message)
		}
		return c.Status(200).JSON("Delegated")

	} else {
		db.
			Where("person_id = ?", c.Params("item_id")).
			Find(&checks)
	}
	return c.Status(404).JSON(&checks)
}

func PatchCheck(c *fiber.Ctx) error {
	db := database.OpenDb()
	var check models.Check

	if c.Params("action") == "create" {
		tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

		itemIDStr := c.Params("item_id")
		itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		check.PersonID = uint(itemID)
		check.Officer = tokenMeta.FullName
		db.Create(&check)

	} else {
		newCheck := models.Check{}
		err := c.BodyParser(&newCheck)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		db.First(&check, c.Params("item_id"))
		db.
			Model(&check).
			Where("item_id = ?", c.Params("item_id")).
			Updates(&newCheck)

		var person models.Person
		db.First(&person, check.PersonID)

		model := models.Conclusion{}
		if check.ConclusionID == model.GetID(utils.Conclusions["saved"]) {
			person.StatusID = model.GetID(utils.Statuses["save"])
		} else if check.ConclusionID == model.GetID(utils.Conclusions["pfo"]) {
			person.StatusID = model.GetID(utils.Statuses["poligraf"])
		} else {
			person.StatusID = model.GetID(utils.Statuses["finish"])
		}
		db.Save(&person)
	}
	return c.Status(200).JSON("Updated")
}

func DeleteCheck(c *fiber.Ctx) error {
	db := database.OpenDb()
	var check models.Check

	db.First(&check, c.Params("item_id"))

	var person models.Person
	db.First(&person, check.PersonID)

	db.Delete(&check)

	var status models.Status
	person.StatusID = status.GetID(utils.Statuses["update"])
	db.Save(&person)

	return c.Status(204).JSON("Check deleted")
}

// Robot Handlers
func GetRobot(c *fiber.Ctx) error {
	db := database.OpenDb()
	var robots []models.Robot
	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&robots)
	return c.Status(200).JSON(robots)
}

func DeleteRobot(c *fiber.Ctx) error {
	db := database.OpenDb()
	var robots []models.Robot
	db.First(&robots, c.Params("item_id"))
	return c.Status(200).JSON("Deleted")
}

func PostRobot(c *fiber.Ctx) error {
	db := database.OpenDb()
	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	cand := models.Person{}
	db.First(&cand, c.Params("item_id"))

	message := models.Message{}

	status := models.Status{}
	if cand.StatusID == status.GetID(utils.Statuses["robot"]) {
		var robot models.Robot

		err := c.BodyParser(&robot)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		persIdStr := c.Params("item_id")
		persId, err := strconv.ParseUint(persIdStr, 10, 64)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		robot.PersonID = uint(persId)
		db.Create(&robot)

		robotPath := "/robots/" + cand.FullName + time.Now().Format("2006-01-02")
		stat, err := os.Stat(robotPath)
		if err == nil {
			if stat.IsDir() {
				candRobotPath := filepath.Join(cand.PathToDocs, "robots")
				_, err := os.Stat(candRobotPath)
				if os.IsNotExist(err) {
					os.Mkdir(candRobotPath, os.ModePerm)
				}
				CopyOrMoveDir(robotPath, cand.PathToDocs, "move")
			}
		}

		message.MessageContent = "Автоматическая проверка кандидата " + cand.FullName + " окончена"
		message.UserID = tokenMeta.UserID
		db.Create(&message)

		cand.StatusID = status.GetID(utils.Statuses["reply"])
		db.Save(&cand)

		return c.Status(200).JSON("Created")
	} else {
		message.MessageContent = "Результат проверки {candidate.fullname} не может быть записан"
		message.UserID = tokenMeta.UserID
		db.Create(&message)
		return c.Status(404).JSON("Not found")
	}
}

// Investigations Handlers
func GetInvestigation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var investigations []models.Investigation
	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&investigations)
	return c.Status(200).JSON(investigations)
}

func PostInvestigation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var investigation models.Investigation
	err := c.BodyParser(&investigation)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	investigation.PersonID = uint(itemID)
	investigation.Officer = tokenMeta.FullName

	db.Create(&investigation)
	return c.Status(200).JSON("Created")
}

func PatchInvestigation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var investigation models.Investigation
	db.First(&investigation, c.Params("item_id"))
	newInvestigation := models.Investigation{}
	err := c.BodyParser(&newInvestigation)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	db.
		Model(&investigation).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&newInvestigation)
	return c.Status(200).JSON("Updated")
}

func DeleteInvestigation(c *fiber.Ctx) error {
	db := database.OpenDb()
	var investigation models.Investigation
	db.First(&investigation, c.Params("item_id"))
	return c.Status(200).JSON("Deleted")
}

// Poligraf Handlers
func GetPoligraf(c *fiber.Ctx) error {
	db := database.OpenDb()
	var poligraf []models.Poligraf
	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&poligraf)
	return c.Status(200).JSON(poligraf)
}

func PostPoligraf(c *fiber.Ctx) error {
	db := database.OpenDb()
	var poligraf models.Poligraf
	err := c.BodyParser(&poligraf)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	poligraf.PersonID = uint(itemID)
	poligraf.Officer = tokenMeta.FullName

	db.Create(&poligraf)

	var person models.Person
	db.First(&person, poligraf.PersonID)

	var status models.Status
	if person.StatusID == status.GetID(utils.Statuses["poligraf"]) {
		person.StatusID = status.GetID(utils.Statuses["finish"])
		db.Save(&person)
	}
	return c.Status(200).JSON("Created")
}

func PatchPoligraf(c *fiber.Ctx) error {
	db := database.OpenDb()
	var poligraf models.Poligraf
	db.First(&poligraf, c.Params("item_id"))
	newPoligraf := models.Poligraf{}
	err := c.BodyParser(&newPoligraf)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	db.
		Model(&poligraf).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&newPoligraf)
	return c.Status(200).JSON("Updated")
}

func DeletePoligraf(c *fiber.Ctx) error {
	db := database.OpenDb()
	var poligraf models.Poligraf
	db.First(&poligraf, c.Params("item_id"))
	return c.Status(200).JSON("Deleted")
}

// Inquiry Handlers
func GetInquiry(c *fiber.Ctx) error {
	db := database.OpenDb()
	var inquiry []models.Inquiry
	db.
		Where("person_id = ?", c.Params("item_id")).
		Find(&inquiry)
	return c.Status(200).JSON(inquiry)
}

func PostInquiry(c *fiber.Ctx) error {
	db := database.OpenDb()
	var inquiry models.Inquiry
	err := c.BodyParser(&inquiry)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	itemIDStr := c.Params("item_id")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	inquiry.PersonID = uint(itemID)
	inquiry.Officer = tokenMeta.FullName

	db.Create(&inquiry)

	return c.Status(200).JSON("Created")
}

func PatchInquiry(c *fiber.Ctx) error {
	db := database.OpenDb()
	var inquiry models.Inquiry
	db.First(&inquiry, c.Params("item_id"))
	newInquiry := models.Inquiry{}
	err := c.BodyParser(&newInquiry)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	db.
		Model(&inquiry).
		Where("item_id = ?", c.Params("item_id")).
		Updates(&newInquiry)
	return c.Status(200).JSON("Updated")
}

func DeleteInquiry(c *fiber.Ctx) error {
	db := database.OpenDb()
	var inquiry models.Inquiry
	db.First(&inquiry, c.Params("item_id"))
	return c.Status(200).JSON("Deleted")
}

// Info Handler
func PostInformation(c *fiber.Ctx) error {
	type ResponseInformationBody struct {
		RegionID uint   `json:"region_id"`
		Start    string `json:"start"`
		End      string `json:"end"`
	}
	var information ResponseInformationBody
	err := c.BodyParser(&information)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	var count int64

	db := database.OpenDb()
	db.
		Table("checks").
		Joins("JOIN persons ON checks.person_id = persons.id").
		Group("checks.conclusion").
		Where("persons.region_id = ?", information.RegionID).
		Where("checks.created_at BETWEEN ? AND ?", information.Start, information.End).
		Count(&count)

	return c.Status(200).JSON(count)
}

// File Handlers
func GetFile(c *fiber.Ctx) error {
	db := database.OpenDb()
	var person models.Person
	db.First(&person, c.Params("item_id"))
	filePath := filepath.Join(os.Getenv("BASE_PATH"), person.PathToDocs, "images", "image.jpg")
	stat, err := os.Stat(filePath)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	if !stat.IsDir() {
		return c.Status(200).SendFile(filePath)
	}
	return c.Status(500).JSON(err)
}

func PostFile(c *fiber.Ctx) error {
	var person models.Person

	db := database.OpenDb()

	if c.Params("action") == "anketa" {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(500).JSON(err)
		}

		tempPath := filepath.Join(os.Getenv("BASE_PATH"), fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02 15-04-05"), file.Filename))
		if err := c.SaveFile(file, tempPath); err != nil {
			return c.Status(500).JSON(err)
		}

		anketa := utils.JsonParse(tempPath)

		db.
			Where("fullname = ?, birthday = ?", anketa.Resume["fullname"], anketa.Resume["birthday"]).
			First(&person)

		if person.ID == 0 {
			db.Table("people").Create(&anketa.Resume)
			person.StatusID = models.Status{}.GetID(utils.Statuses["new"])

		} else {
			person.StatusID = models.Status{}.GetID(utils.Statuses["update"])
			db.
				Table("people").
				Where("id = ?", person.ID).
				Updates(&anketa.Resume)
		}

		anketa.Staff["person_id"] = string(rune(person.ID))
		db.Table("staff").Create(&anketa.Staff)

		anketa.Document["person_id"] = string(rune(person.ID))
		db.Table("documents").Create(&anketa.Document)

		for _, address := range anketa.Addresses {
			address["person_id"] = string(rune(person.ID))
			db.Table("addresses").Create(&address)
		}

		for _, workplace := range anketa.Workplaces {
			workplace["person_id"] = string(rune(person.ID))
			db.Table("workplaces").Create(&workplace)
		}

		for _, contact := range anketa.Contacts {
			contact["person_id"] = string(rune(person.ID))
			db.Table("contacts").Create(&contact)
		}

		for _, affiliation := range anketa.Affilations {
			affiliation["person_id"] = string(rune(person.ID))
			db.Table("affilations").Create(&affiliation)
		}

		person.PathToDocs = makeFolder(person.FullName, person.ID)
		actionPath := filepath.Join(os.Getenv("BASE_PATH"), person.PathToDocs, c.Params("action"))
		err = os.Mkdir(actionPath, 0755)
		if err == nil {
			os.Rename(tempPath, filepath.Join(actionPath, file.Filename))
		}
		db.Save(&person)

	} else {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(500).JSON(err)
		}
		files := form.File["files"]

		db.Table("people").Where("id = ?", c.Params("item_id")).First(&person)
		actionPath := filepath.Join(os.Getenv("BASE_PATH"), person.PathToDocs, c.Params("action"))
		err = os.Mkdir(actionPath, 0755)
		if err != nil {
			return c.Status(500).JSON(err)
		}

		if c.Params("action") == "image" {
			if err := c.SaveFile(files[0], filepath.Join(actionPath, "image.jpg")); err != nil {
				return c.Status(500).JSON(err)
			}

		} else {
			for _, file := range files {
				if err := c.SaveFile(file, filepath.Join(actionPath, file.Filename)); err != nil {
					return c.Status(500).JSON(err)
				}
			}
		}
	}

	return c.Status(200).JSON(person.ID)
}
