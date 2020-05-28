package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"sprint3/Helpers"
	"sprint3/Models"
	"text/template"
)

func GetCompany() {
	db := Helpers.DbConnect()
	err := db.Debug().Model(&Models.Driver{}).Find(&drivers).Error
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}

var Company = []Models.Driver{}

// Get Company Template Template Handler
// @Summary Driver Late
// @Description Send To Company when Driver is late
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /colate  [get]
func CoAlert_DriverLate(w http.ResponseWriter, r *http.Request) {
	GetCompany()
	tpl := template.Must(template.ParseFiles("templates/company/CoAlert_DriverLate.tmpl"))
	fmt.Println(Company)
	err2 := tpl.Execute(w, Company)
	if err2 != nil {
		log.Fatal("Error executing template:", err2)
	}

}

// Get Company Template Handler
// @Summary No Driver
// @Description Send To Company When there is no driver
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /cono  [get]
func CoNoDriver(w http.ResponseWriter, r *http.Request) {
	GetCompany()
	tpl := template.Must(template.ParseFiles("templates/company/CoNoDriver.tmpl"))
	err2 := tpl.Execute(w, Company)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Company Template Template Handler
// @Summary  Cancelled delivery
// @Description Send To Company When Cancelled their Delivery
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /conodelivery [get]
func CoCancelDelivery(w http.ResponseWriter, r *http.Request) {
	GetCompany()
	tpl := template.Must(template.ParseFiles("templates/company/CoCancelDelivery.tmpl"))
	err2 := tpl.Execute(w, Company)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Company Template Handler
// @Summary  Ride cancellation
// @Description Send To Company When their ride is Cancelled
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /conoride  [get]
func CoCancelRide(w http.ResponseWriter, r *http.Request) {
	GetCompany()
	tpl := template.Must(template.ParseFiles("templates/company/CoCancelRide.tmpl"))
	err2 := tpl.Execute(w, Company)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Company Template Handler
// @Summary Delivery Confirmation
// @Description Send to Company When Their Delivery is Confirmed
// @Tags Templates
// @Accept json
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /coyesdelivery  [get]
func CoConfirmDelivery(w http.ResponseWriter, r *http.Request) {
	GetCompany()
	tpl := template.Must(template.ParseFiles("templates/company/CoConfirmDelivery.tmpl"))
	err2 := tpl.Execute(w, Company)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get Company Template Handler
// @Summary Ride confirmation
// @Description Send To Company When Ride is Confirmed
// @Tags Templates
// @Accept json
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /cokride [get]
func CoConfirmRide(w http.ResponseWriter, r *http.Request) {
	GetCompany()
	tpl := template.Must(template.ParseFiles("templates/company/CoConfirmRide.tmpl"))
	err2 := tpl.Execute(w, Company)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}
