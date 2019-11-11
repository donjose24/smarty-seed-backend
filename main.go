package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmramos02/smarty-seed-backend/api"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/config"
)

func main() {
	//migrate Databases first, if needs any updating.
	db, err := gorm.Open("postgres", config.GetDatabaseUrl())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err.Error()))
	}
	db.AutoMigrate(&models.User{})
	router := api.Initialize(db)

	router.Run()
}
