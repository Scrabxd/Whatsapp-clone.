package Server

import (
	controllers "WhatsAppClone/Controllers"
	db "WhatsAppClone/DB"
	"WhatsAppClone/Helpers"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Server() {
	// Testing and connection to database
	err := db.MongoDb()
	if err != nil {
		log.Fatal(err)
	}
	// Creation of the server
	app := fiber.New()

	// Setting the port to be used
	port := Helpers.GetEnv("port")
	if port == "" {
		port = "5000"
	}

	//Setting cors
	app.Use(cors.New(cors.ConfigDefault))

	// Routing

	//User
	app.Get("/", controllers.HelloWold)
	app.Post("/CreateUser", controllers.CreateUser)
	app.Delete("/DeleteUser", controllers.DeleteUser)

	//Starting server
	log.Fatal(app.Listen(":" + port))
}
