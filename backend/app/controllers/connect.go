package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/utils"
	"backend/platform/database"
)

func GetConnects(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{}, []string{})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	db := database.OpenDb()
	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}
	var pagination = 16
	var hasPrev, hasNext bool
	var connects []models.Connection

	searchData := struct {
		Search string `json:"search"`
	}{}
	if err := c.BodyParser(&searchData); err != nil {
		log.Println(err)
	}
	result := []interface{}{}

	companies := db.
		Select("company").
		Find(&connects)
	cities := db.
		Select("city").
		Find(&connects)

	query := db.
		Find(&connects).
		Limit(pagination).
		Offset(pagination * (intPage - 1))

	if searchData.Search != "" {
		query = query.Where("id = ?", searchData.Search)
	}

	query.Find(&result)

	if intPage > 1 {
		hasPrev = true
	}
	if len(result) == pagination {
		hasNext = true
	}
	return c.JSON(fiber.Map{
		"result":    result,
		"hasNext":   hasNext,
		"hasPrev":   hasPrev,
		"companies": companies,
		"cities":    cities,
	})
}
