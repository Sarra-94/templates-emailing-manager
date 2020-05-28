package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sprint3/Helpers"
	"sprint3/Models"
	"text/template"

	"github.com/gorilla/mux"
)

// get all tempaltes (depends on the component name and saas offices)
func GetAllTemplates(w http.ResponseWriter, r *http.Request) {
	// get the name of saas office & company from url
	paramatres := mux.Vars(r)
	fmt.Println("parameters are", paramatres)
	// get template's name from body
	body, err := ioutil.ReadAll(r.Body)
	var data Data
	err = json.Unmarshal(body, &data)

	fmt.Println("body is", data)
	if err != nil {
		fmt.Println(err)
	}

	file, err1 := os.Open("templates")
	if err1 != nil {
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

func Alert_DriverLate(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/Alert_DriverLate.tmpl"))
	fmt.Println(drivers)
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template:", err2)
	}

}

func Getdrivers() {
	db := Helpers.DbConnect()
	err := db.Debug().Model(&Models.Driver{}).Find(&drivers).Error
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}

var drivers = []Models.Driver{}
