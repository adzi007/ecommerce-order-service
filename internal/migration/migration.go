package main

import (
	"github.com/adzi007/ecommerce-order-service/config"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-order-service/internal/model"
)

func main() {
	config.LoadConfig()

	db := database.NewPostgreesDatabase()

	appDbMigrate(db)
}

func appDbMigrate(db database.Database) {

	// db.GetDb().Migrator().CreateTable(&entity.Cart{})

	// err := db.GetDb().Migrator().AutoMigrate(&model.Order{}, &model.OrderDetail{})
	err := db.GetDb().Migrator().CreateTable(&model.Order{}, &model.OrderDetail{})

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

}
