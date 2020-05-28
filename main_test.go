package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	if DB, err = OpenTestConnection(); err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
	}

}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")

	fmt.Println("testing postgres...")
	if dbDSN == "" {
		dbDSN = "user=postgres password=root DB.name=craft port=5432 sslmode=disable"
	}
	db, err = gorm.Open("postgres", dbDSN)

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.LogMode(true)
	} else if debug == "false" {
		db.LogMode(false)
	}

	db.DB().SetMaxIdleConns(10)

	return
}

type Driver struct {
	Driver_full_name string
	Request_id       string
}

func TestRow(t *testing.T) {
	user1 := Driver{Driver_full_name: "Ahmed", Request_id: "AH14"}
	user2 := Driver{Driver_full_name: "Ahmed", Request_id: "AH14"}
	user3 := Driver{Driver_full_name: "RowUser3", Request_id: "time25"}
	DB.Save(&user1).Save(&user2).Save(&user3)

	row := DB.Table("users").Where("driver_full_name = ?", user2.Driver_full_name).Select("request_id").Row()
	var request_id, AH14 string
	row.Scan(&request_id)
	if request_id != AH14 {
		t.Errorf("Scan with Row")
	}
}
