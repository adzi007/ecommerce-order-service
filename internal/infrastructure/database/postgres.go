package database

import (
	"fmt"
	"time"

	"github.com/adzi007/ecommerce-order-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	// once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgreesDatabase() Database {

	// once.Do(func() {

	dbUsername := config.ENV.DB_USERNAME
	dbPassword := config.ENV.DB_PASSWORD
	dbName := config.ENV.DB_NAME
	dbHost := config.ENV.DB_HOST
	dbPort := config.ENV.DB_PORT

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// dsn := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUsername, dbPassword, dbName, dbPort)

	var db *gorm.DB
	var err error

	// Retry connecting to the database up to 10 times
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
		})
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to the database. Retrying in 2 seconds... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic("failed to connect database after 10 attempts")
	}

	dbInstance = &postgresDatabase{Db: db}

	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
