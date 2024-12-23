package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ayoadeoye1/insta-shop-screening/helper"
)

func ConnectDB() *gorm.DB {
	host := EnvFile["DB_HOST"]
	port := EnvFile["DB_PORT"]
	user := EnvFile["DB_USER"]
	password := EnvFile["DB_PASSWORD"]
	dbName := EnvFile["DBNAME"]

	dbConfig := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("DB Connection Successful")

	return db
}
