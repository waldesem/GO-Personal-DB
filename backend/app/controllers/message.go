package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/middlewares"
	"backend/platform/database"
)

func GetMessages(c *fiber.Ctx) error {
	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}

	db := database.OpenDb()
	var messages []models.Message
	var pagination = 16
	var hasPrev, hasNext bool

	switch c.Params("action") {
	case "new":
		db.
			Where("status = ?, user_id = ?", "new", tokenMeta.UserID).
			Find(&messages).
			Limit(pagination).
			Offset(pagination * (intPage - 1))
	case "read":
		db.
			Where("status = ?, user_id = ?", "new", tokenMeta.UserID).
			Find(&messages)
		for message := range messages {
			messages[message].StatusRead = "reply"
			db.Save(&messages[message])
		}
		db.
			Where("status = ?, user_id = ?", "new", tokenMeta.UserID).
			Find(&messages).
			Limit(pagination).
			Offset(pagination * (intPage - 1))
	}

	if intPage > 1 {
		hasPrev = true
	}
	if len(messages) == pagination {
		hasNext = true
	}
	result, err := json.Marshal(messages)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{"result": result, "hasNext": hasNext, "hasPrev": hasPrev})
}

func DeleteMessage(c *fiber.Ctx) error {
	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)

	db := database.OpenDb()
	var messages []models.Message

	db.
		Find(&messages).
		Where("person_id = ?", tokenMeta.UserID).
		Delete(&messages)
	return c.SendStatus(fiber.StatusOK)
}
