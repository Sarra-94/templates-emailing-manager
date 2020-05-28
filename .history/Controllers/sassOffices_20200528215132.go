package Controllers

import (
	"encoding/json"
	"net/http"
	"sprint3/Helpers"

	"github.com/gorilla/mux"
)

// Get Backend user Template Handler
// @Summary Permission for acces
// @Description Send To Backend user
// @Tags Templates
// @Accept json
// @Produce json
// @success 200 {object} Models.Driver
// @Router /us  [get]
func GetOfficesBySaasCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	saasParam := params["saas_company"]
	db := Helpers.DbConnect()
	var saasCompany Models.SaasCompany
	db.Model(&saasCompany).Where("name = ?", saasParam).Find(&saasCompany)
	var saasOffice []Models.SaasOffice
	db.Model(&saasOffice).Where("saas_company_id = ?", saasCompany.ID).Find(&saasOffice)

	json.NewEncoder(w).Encode(saasOffice)

}
