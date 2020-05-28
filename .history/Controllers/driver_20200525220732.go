package Controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sprint3/Helpers"
	"sprint3/Models"
	"text/template"
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

// Get Notway driver Template Handler
// @Summary  not way driver template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /noy  [get]
func NotwDriver(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/NotwayDriver.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get alert driver late Template Handler
// @Summary  No driver template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /no  [get]
func NoDriver(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/NoDriver.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Cancel delivery Template Handler
// @Summary  Cancel delivery template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /nodelivery [get]
func CancelDelivery(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/CancelDelivery.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Cancel client  ride Template Handler
// @Summary  client cancel ride template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /noclient  [get]
func CancelClientRide(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/CancelClientRide.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Cancel ride Template Handler
// @Summary  Ride canceled Template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /noride  [get]
func CancelRide(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/CancelRide.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Confirm delivery Template Handler
// @Summary  Confirmed Delivery Template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /yesdelivery  [get]
func ConfirmDelivery(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/ConfirmDelivery.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Confirm client ride Template Handler
// @Summary  Ride confirmed by client template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router  /yesrd [get]
func ConfirmClientRide(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/ConfirmClientRide.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Confirm ride Template Handler
// @Summary  Confirmed Ride template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /okride  [get]
func ConfirmRide(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/ConfirmRide.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Driver Invoice Template Handler
// @Summary Driver Invoice Template
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept json
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /invoice  [get]
func DriverInvoice(w http.ResponseWriter, r *http.Request) {
	Getdrivers()
	tpl := template.Must(template.ParseFiles("templates/driver/DriverInvoice.tmpl"))
	err2 := tpl.Execute(w, drivers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

func Fil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path string
	fmt.Println("path for template :")
	fmt.Scanf("%s", &path)
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		json.NewEncoder(w).Encode(eachline)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(eachline))
	}

}
