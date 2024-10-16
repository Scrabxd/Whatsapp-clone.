package db

import (
	"WhatsCl/Helpers"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgreSQL() (*gorm.DB, error) {

	var DSN = Helpers.GetEnv("DSN_LOCAL")

	if strings.ToUpper(Helpers.GetEnv("ENV")) == "DEPLOY" {
		DSN = Helpers.GetEnv("DSN_DEPLOY")
	}

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
