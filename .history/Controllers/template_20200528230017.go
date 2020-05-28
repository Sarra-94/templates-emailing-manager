package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Data struct {
	Name string
}

// get all tempaltes (depends on the component name and saas offices)
func GetAllTemplates(w http.ResponseWriter, r *http.Request) {
	// get the name of saas office & company from url
	paramatres := mux.Vars(r)
	saasComp := paramatres["saascompany"]
	saasoff := paramatres["saasoffice"]
	fmt.Println("parameters are", paramatres)
	// get component's from body
	body, err := ioutil.ReadAll(r.Body)
	var data Data
	err = json.Unmarshal(body, &data)
	fmt.Println("body is", data)
	if err != nil {
		fmt.Println(err)
	}
	str := "/" + saasComp + "/" + saasoff + "/" + data.Name
	// /get all templates name
	file, err1 := os.Open("templates" + str)
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

// fonction get contenu du template selon saascompany saas offices and component
func Alert_DriverLate(w http.ResponseWriter, r *http.Request) {
	// get the name of saas office & company from url
	paramatres := mux.Vars(r)
	saasComp := paramatres["saascompany"]
	saasoff := paramatres["saasoffice"]
	comp := paramatres["component"]
	fmt.Println("parameters are", paramatres)
	// get component's from body
	body, err := ioutil.ReadAll(r.Body)
	var data Data
	err = json.Unmarshal(body, &data)
	fmt.Println("body is", data)
	str := "/" + saasComp + "/" + saasoff + "/" + comp + "/" + data.Name
	tpl := template.Must(template.ParseFiles("templates" + str))
	err2 := tpl.Execute(w, nil)
	if err2 != nil {
		log.Fatal("Error executing template:", err2)
	}

}
