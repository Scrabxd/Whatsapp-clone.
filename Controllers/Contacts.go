package controllers

import (
	db "WhatsCl/DB"
	models "WhatsCl/Models/postgres"

	"github.com/gofiber/fiber/v2"
)

var database, _ = db.PostgreSQL()

func CreateContactCellphone(c *fiber.Ctx) error {

	name, last_name, phone_number := c.FormValue("name"), c.FormValue("last_name"), c.FormValue("phone_number")

	contact := models.Contact{Name: name, Last_name: last_name, Phone_Number: phone_number}
	result := database.FirstOrCreate(&contact, models.Contact{Phone_Number: phone_number})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Error while finding or creating data.",
			"Error":   result.Error,
			"Stauts":  500,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Message": "Success Creating register.",
		"Data":    contact,
		"Status":  200,
	})
}

func RetrieveAllContacts(c *fiber.Ctx) error {

	var users []models.Contact

	result := database.Find(&users)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Error Fetching contacts",
			"Error":   result.Error,
			"Stauts":  500,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Message": "All contacts retrieved",
		"Data":    users,
		"Status":  200,
	})
}
