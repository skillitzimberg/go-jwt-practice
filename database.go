package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func GetDatabase() *gorm.DB {
	// host := "localhost"
	// port := 5432
	user := "foreignfood"
	password := ""
	dbname := "foreignfood"
	url := fmt.Sprintf("%s:%s@tcp(127.0.0.1:5432)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
	connection, err := gorm.Open("psql", url)
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}
func InitialMigration() {
	connection := GetDatabase()
	defer Closedatabase(connection)
	connection.AutoMigrate(User{})
}

func Closedatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}