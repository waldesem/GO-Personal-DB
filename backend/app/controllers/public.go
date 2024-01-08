package controllers

import (
	"encoding/json"
	"log"
	"strconv"

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
	auth, err := utils.RolesGroupsInToken(c, []string{"user"}, []string{"staffsec"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

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
			Offset(pagination * (intPage - 1))
	case "officer":
		var checks []models.Check
		statusFin := models.Status{}.GetID(utils.Statuses["finish"])
		statusCan := models.Status{}.GetID(utils.Statuses["cancel"])
		db.
			Find(&persons).
			Where("status_id != ? OR status_id != ?", statusFin, statusCan).
			Limit(pagination).
			Offset(pagination * (intPage - 1)).
			Joins("JOIN checks ON checks.person_id = people.id").
			Where(models.Check{Officer: "current"}).
			Find(&checks)
	case "search":
		db.Where("fullname LIKE ?", "%"+payload.Search+"%").Find(&persons).Limit(10).Offset(10 * (intPage - 1))
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
