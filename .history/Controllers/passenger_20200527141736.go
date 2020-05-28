package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"sprint3/Helpers"
	"sprint3/Models"
	"text/template"
)

var customers = []Models.Driver{}

func Getcustomers() {
	db := Helpers.DbConnect()
	err := db.Debug().Model(&Models.Driver{}).Find(&customers).Error
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}

// Get Alert customer Template Handler
// @Summary Customer Late
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /latecustomer  [get]
func Alert_CustomerLate(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/Alert_CustomerLate.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get No customer Template Handler
// @Summary No customer
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /nocustomer  [get]
func NoCustomer(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/NoCustomer.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get customer cancel delivery Template Handler
// @Summary Delivey Cancelled
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /nodeliveryc [get]
func CancelDeliveryCustomer(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/CancelDeliveryCustomer.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get A customer cancel Driver ride Template Handler
// @Summary Driver Cancel Client Ride
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /nod  [get]
func CancelDriverRide(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/CancelDRide.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get customer cancel ride Template Handler
// @Summary Driver  Cancel Ride
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /noridec  [get]
func CustomerCancelRide(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/CustomerCancelRide.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}
}

// Get  customer confirm delivery Template Handler
// @Summary Driver Confirm Delivery
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /yesdelic  [get]
func CustomerConfirmDelivery(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/CustomerConfirmDelivery.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get  customer confirm ride Template Handler
// @Summary Driver Confirm Ride
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /okridec  [get]
func ConfirmDriverRide(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/ConfirmDriverRide.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// Get customer Invoice Template Handler
// @Summary Customer Invoice
// @Description generating template  which will be sended to driver
// @Tags Templates
// @Accept text/html
// @Produce text/html
// @success 200 {object} Models.Driver
// @Router /Cinvoice [get]
func CustomerInvoice(w http.ResponseWriter, r *http.Request) {
	Getcustomers()
	tpl := template.Must(template.ParseFiles("templates/passenger/CustomerInvoice.tmpl"))
	err2 := tpl.Execute(w, customers)
	if err2 != nil {
		log.Fatal("Error executing template: ", err2)
	}

}

// get all drivers tempaltes
func ConfirmPassengerRide(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("templates/driver")
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
