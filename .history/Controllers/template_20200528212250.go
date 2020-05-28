package Controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sprint3/Helpers"
	"sprint3/Models"
	"text/template"

	"github.com/gorilla/mux"
)

func Getdrivers() {
	db := Helpers.DbConnect()
	err := db.Debug().Model(&Models.Driver{}).Find(&drivers).Error
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}

var drivers = []Models.Driver{}

// get all drivers tempaltes
func GetAllTemplates(w http.ResponseWriter, r *http.Request) {

	paramatres := mux.Vars(r)

	file, err := os.Open("templates/saasoffice/nomcomponent")
	if err != nil {
		log.Fatalf("failed opening drivers templates: %s", err)
	}
	defer file.Close()
	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	var lists []string
	for _, name := range list {
		lists = append(lists, name)
	}
	json.NewEncoder(w).Encode(lists)
}

// Get alert driver late Template Handler
// @Summary  driver alert late  template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /late  [get]
func Alert_DriverLate(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/Alert_DriverLate.tmpl"))
	fmt.Println(drivers)
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template:", err2)
	}

}
