package Helpers

import (
	"fmt"
	"log"
	"os"
	"sprint3/Models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetVarFromEnv() string {
	host := os.Getenv("host")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("port")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("user")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("password")
	if password == "" {
		password = "root"
	}
	dbname := os.Getenv("name")
	if dbname == "" {
		dbname = "yuso"
	}
	if os.Getenv("DATABASE_URL") != "" {
		ConnectionString := os.Getenv("DATABASE_URL")
		return ConnectionString
	}
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return ConnectionString
}

func DbConnect() *gorm.DB {
	connection := GetVarFromEnv()
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migration() {
	db := DbConnect()
	db.AutoMigrate(&Models.Driver{})

}
