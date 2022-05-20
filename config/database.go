package config

import (
	"base-api/app/entities"
	"base-api/routes"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) {

	DBURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)

	DB, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect to database")
		panic(err.Error())
	} else {
		fmt.Println("We are connected to the database")
	}

	DB.AutoMigrate(
		&entities.Participant{},
	)

	routes.InitializeRoutes(DB)
}
