package Server

import (
	controllers "WhatsCl/Controllers"
	db "WhatsCl/DB"
	"WhatsCl/Helpers"
	sockets "WhatsCl/Sockets"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func Server() {
	// Testing and connection to MONGODB
	err := db.MongoDb()
	if err != nil {
		log.Fatal(err)
	}

	//TESTING AND CONNECTION TO POSTGRES
	_, err = db.PostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to PG stablished correctly ✅")
	// Creation of the server
	app := fiber.New()

	// Setting the port to be used
	port := Helpers.GetEnv("port")
	if port == "" {
		port = "5000"
	}

	//Setting cors
	app.Use(cors.New(cors.ConfigDefault))

	//User
	app.Get("/", controllers.HelloWold)
	app.Post("/CreateUser", controllers.CreateUser)
	app.Post("/SearchUsers", controllers.SearchUsers)
	app.Delete("/DeleteUser", controllers.DeleteUser)
	app.Put("/UpdatUser", controllers.UpdateData)

	//Socket route
	app.Get("/ws", websocket.New(sockets.HandleWebSocket))

	//Cellphone simulation
	app.Get("/Contacts", controllers.RetrieveAllContacts)
	app.Post("/CreateContact", controllers.CreateContactCellphone)
	//Starting server
	log.Fatal(app.Listen(":" + port))
}
