package Routers

import (
	"net/http"
	Controllers "sprint3/Controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	mux := mux.NewRouter()
	// Driver router
	mux.HandleFunc("/{saas_company}/offices", Controllers.GetOfficesBySaasCompany).Methods("GET")
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, to Driver section ."))
	})
	// Passenger router
	// mux.HandleFunc("/latecustomer", Controllers.Alert_CustomerLate).Methods("GET")

	return mux
}
