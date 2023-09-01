package models

import (
	"fmt"
	"golang-rest-api-template/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db_hostname := env.PostgresHost()
	db_name := env.PostgresDBName()
	db_user := env.PostgresUser()
	db_pass := env.PostgresPassword()
	db_port := env.PostgresPort()

	dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)
	database, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{})
	database.AutoMigrate(&User{})

	DB = database
}
