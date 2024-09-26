package controllers

import (
	models "WhatsAppClone/Models"
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c fiber.Ctx) error {

	var res = make(map[string]interface{})

	extension, phone_number, country := c.FormValue("extension"), c.FormValue("phone_number"), c.FormValue("country")

	if extension == "" || phone_number == "" || country == "" {
		res = map[string]interface{}{
			"Message": "Error, one or more parameters are empty",
			"Status":  400,
		}
		return c.Status(400).JSON(res)
	}

	user := models.NewUser(extension, phone_number, country)

	err := mgm.Coll(user).Create(user)
	if err != nil {
		res = map[string]interface{}{
			"Message": "Error while creating User",
			"Status":  500,
			"Error":   err,
		}
		return c.Status(500).JSON(res)
	}

	res = map[string]interface{}{
		"Message": "User Created Successfully",
		"Status":  200,
		"Data":    user,
	}

	return c.Status(500).JSON(res)
}

func DeleteUser(c fiber.Ctx) error {
	var res = make(map[string]interface{})

	phone_number, extension := c.FormValue("phone_number"), c.FormValue("extension")

	filter := bson.M{"Phone_Number": phone_number, "Extension": extension}

	data, err := mgm.Coll(&models.Users{}).DeleteOne(context.Background(), filter)
	if err != nil {
		res = map[string]interface{}{
			"Message": "Couldn't delete User",
			"Status":  500,
			"Error":   err,
		}
		return c.Status(500).JSON(res)
	}

	res = map[string]interface{}{
		"Message":       "User deleted Sucessfully",
		"Status":        200,
		"Phone_deleted": phone_number,
		"Data":          data,
	}
	return c.Status(200).JSON(res)
}
