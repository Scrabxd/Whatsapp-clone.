package db

import (
	"WhatsCl/Helpers"
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDb() error {
	uri := Helpers.GetEnv("MONGO_URI")
	if uri == "" {
		log.Println("MONGO_URI environment variable is not set.")
		return fmt.Errorf("MONGO_URI is required")
	}

	err := mgm.SetDefaultConfig(nil, "WhatsCl", options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("CONNECTION WITH MONGO REJECTED ❌: %v\n", err)
		return err
	}

	log.Println("CONNECTION WITH MONGO ESTABLISHED ✅")
	return nil
}
