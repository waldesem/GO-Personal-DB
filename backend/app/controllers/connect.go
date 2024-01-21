package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/platform/database"
)

func GetConnects(c *fiber.Ctx) error {
	db := database.OpenDb()
	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}
	var pagination = 16
	var hasPrev, hasNext bool
	var connects, companies, cities []models.Connection

	searchData := struct {
		Search string `json:"search"`
	}{}
	if err := c.BodyParser(&searchData); err != nil {
		log.Println(err)
	}

	db.
		Select("company").
		Find(&companies)
	db.
		Select("city").
		Find(&cities)

	query := db.
		Find(&connects).
		Limit(pagination).
		Offset(pagination * (intPage - 1))

	if searchData.Search != "" {
		query = query.Where("company = ?", searchData.Search)
	}
	query.Find(&connects)

	if intPage > 1 {
		hasPrev = true
	}
	if len(connects) == pagination {
		hasNext = true
	}
	return c.JSON(fiber.Map{
		"result":    connects,
		"hasNext":   hasNext,
		"hasPrev":   hasPrev,
		"companies": companies,
		"cities":    cities,
	})
}

func PostConnect(c *fiber.Ctx) error {
	db := database.OpenDb()
	var connect models.Connection

	if err := c.BodyParser(&connect); err != nil {
		return c.Status(500).JSON(err)
	}

	db.Create(&connect)

	return c.Status(200).JSON("Created")
}

func PatchConnect(c *fiber.Ctx) error {
	db := database.OpenDb()
	var connect models.Connection

	if err := c.BodyParser(&connect); err != nil {
		return c.Status(500).JSON(err)
	}

	db.
		Model(&connect).
		Where("id = ?", c.Params("item_id")).
		Updates(&connect)

	return c.Status(200).JSON("Updated")
}

func DeleteConnect(c *fiber.Ctx) error {
	db := database.OpenDb()
	var connect models.Connection

	db.Delete(&connect, c.Params("item_id"))

	return c.Status(200).JSON("Deleted")
}
