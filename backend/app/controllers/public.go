package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
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

			body := struct {
				Person models.Person
				Docs   models.Document
				Addr   models.Address
			}{
				Person: person,
				Docs:   docs,
				Addr:   addr,
			}

			jsonBody, err := json.Marshal(body)
			if err != nil {
				return c.Status(500).JSON(err)
			}

			resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonBody))
			if err != nil {
				return c.Status(500).JSON(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return c.Status(500).JSON(err)
			} else {
				person.StatusID = models.Status{}.GetID(utils.Statuses["robot"])
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
