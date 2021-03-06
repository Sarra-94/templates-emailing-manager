package Routers

import (
	"net/http"
	Controllers "sprint3/Controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	mux := mux.NewRouter()
	// Driver router
	mux.HandleFunc("/{saas_company}/offices", Controllers.Alert_DriverLate).Methods("GET")
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, to Driver section ."))
	})

	// Passenger router
	mux.HandleFunc("/latecustomer", Controllers.Alert_CustomerLate).Methods("GET")
	mux.HandleFunc("/nocustomer", Controllers.NoCustomer).Methods("GET")
	mux.HandleFunc("/nodeliveryc", Controllers.CancelDeliveryCustomer).Methods("GET")
	mux.HandleFunc("/nod", Controllers.CancelDriverRide).Methods("GET")
	mux.HandleFunc("/noridec", Controllers.CustomerCancelRide).Methods("GET")
	mux.HandleFunc("/yesdelic", Controllers.CustomerConfirmDelivery).Methods("GET")
	mux.HandleFunc("/tmplpass", Controllers.ConfirmPassengerRide).Methods("GET")
	mux.HandleFunc("/tmplpassenger", Controllers.ReadCurrentDir).Methods("GET")
	mux.HandleFunc("/Cinvoice", Controllers.CustomerInvoice).Methods("GET")

	// Company router
	mux.HandleFunc("/colate", Controllers.CoAlert_DriverLate).Methods("GET")
	mux.HandleFunc("/conodelivery", Controllers.CoCancelDelivery).Methods("GET")
	mux.HandleFunc("/conoride", Controllers.CoCancelRide).Methods("GET")
	mux.HandleFunc("/coyesdelivery", Controllers.CoConfirmDelivery).Methods("GET")
	mux.HandleFunc("/cokride", Controllers.CoConfirmRide).Methods("GET")
	mux.HandleFunc("/tmplcompany", Controllers.ConfirmCompanyRide).Methods("GET")
	mux.HandleFunc("/cono", Controllers.CoNoDriver).Methods("GET")

	// backuser router
	mux.HandleFunc("/us", Controllers.Access).Methods("GET")
	mux.HandleFunc("/acc", Controllers.Account).Methods("GET")
	mux.HandleFunc("/dnc", Controllers.Dnacces).Methods("GET")

	// files
	mux.HandleFunc("/fex", Controllers.Fexpired).Methods("GET")
	mux.HandleFunc("/fre", Controllers.Frefused).Methods("GET")
	mux.HandleFunc("/fva", Controllers.Fvalidated).Methods("GET")

	//Welcome
	mux.HandleFunc("/bwl", Controllers.Backwelm).Methods("GET")
	mux.HandleFunc("/cwl", Controllers.Cmwelm).Methods("GET")
	mux.HandleFunc("/dwl", Controllers.Drwelm).Methods("GET")
	mux.HandleFunc("/prwl", Controllers.Pawelm).Methods("GET")
	mux.HandleFunc("/pswl", Controllers.Pswelm).Methods("GET")

	return mux

}
