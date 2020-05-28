package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"sprint3/Helpers"
	"sprint3/Models"
	"text/template"
)

func Getbackuser() {
	db := Helpers.DbConnect()
	err := db.Debug().Model(&Models.Driver{}).Find(&backuser).Error
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}

var backuser = []Models.Driver{}

// Get Backend user Template Handler
// @Summary Permission for acces
// @Description Send To Backend user
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /us  [get]
func Access(w http.ResponseWriter, r *http.Request) {
	Getbackuser()
	tpl := template.Must(template.ParseFiles("templates/backuser/Cnfacces.tmpl"))
	fmt.Println(backuser)
	err2 := tpl.Execute(w, backuser)
	if err2 != nil {
		log.Fatal("Error executing template:", err2)
	}

}

// Get Backend user Template Handler
// @Summary Confirm Account
// @Description Send To Backend user
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /acc  [get]
func Account(w http.ResponseWriter, r *http.Request) {
	Getbackuser()
	tpl := template.Must(template.ParseFiles("templates/backuser/Cnfaccount.tmpl"))
	err2 := tpl.Execute(w, backuser)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Backend user Template Handler
// @Summary  No Access
// @Description Send For Backend User
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /dnc [get]
func Dnacces(w http.ResponseWriter, r *http.Request) {
	Getbackuser()
	tpl := template.Must(template.ParseFiles("templates/backuser/Dnacces.tmpl"))
	err2 := tpl.Execute(w, backuser)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}
