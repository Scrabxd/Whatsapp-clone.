package controllers

import (
	"WhatsAppClone/Helpers"
	models "WhatsAppClone/Models"
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c fiber.Ctx) error {
	extension, phone_number, country, username := c.FormValue("extension"), c.FormValue("phone_number"), c.FormValue("country"), c.FormValue("username")

	user := models.NewUser(extension, phone_number, country, username)

	var findUser models.Users

	filter := bson.M{"Phone_Number": phone_number}

	if extension == "" || phone_number == "" || country == "" || username == "" {
		return c.Status(400).JSON(fiber.Map{
			"Message": "Error, one or more parameters are empty",
			"Status":  400})
	}

	err := mgm.Coll(&models.Users{}).FindOne(context.Background(), filter).Decode(&findUser)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"Message": "Error while finding Number.",
			"Status":  500,
			"Error":   err,
		})
	}

	if findUser.Phone_Number != "" {
		return c.Status(400).JSON(fiber.Map{
			"Message": "User already exists.",
			"Status":  400,
		})
	}

	err = mgm.Coll(user).Create(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Error while creating User",
			"Status":  500,
			"Error":   err})
	}

	return c.Status(500).JSON(fiber.Map{
		"Message": "User Created Successfully",
		"Status":  200,
		"Data":    user,
	})
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

func SearchUsers(c fiber.Ctx) error {
	var res = make(map[string]interface{})

	var users []models.Users

	ExistingNumbers := make([]bson.M, 0)

	phone_numbers, err := Helpers.Parser(c.Body())
	if err != nil {
		res = map[string]interface{}{
			"Message": "Error parsing Phone Numbers",
			"Status":  500,
			"Error":   err,
		}

		return c.Status(200).JSON(res)
	}

	err = mgm.Coll(&models.Users{}).SimpleFind(&users, bson.M{"Phone_Number": bson.M{"$in": phone_numbers.Phone_number}})
	if err != nil {
		return err
	}

	for _, user := range users {
		ExistingNumbers = append(ExistingNumbers, bson.M{"Phone_Number": user.Phone_Number, "Username": user.Username})
	}

	res = map[string]interface{}{
		"Message":         "Phone Number check complete",
		"Status":          200,
		"ExistingNumbers": ExistingNumbers,
	}

	return c.Status(200).JSON(res)

}
