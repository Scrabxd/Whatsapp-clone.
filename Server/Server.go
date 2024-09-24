package Server

import (
	controllers "WhatsAppClone/Controllers"
	"WhatsAppClone/DB"
	"WhatsAppClone/Helpers"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Server() {

	defer func() {
		db, _ := DB.ConnectPG()

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.Close()
	}()

	app := fiber.New()

	port := Helpers.GetEnv("port")
	if port == "" {
		port = "5000"
	}

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/", controllers.Main)

	log.Fatal(app.Listen(":" + port))
}
