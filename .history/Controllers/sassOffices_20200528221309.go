package Controllers

import (
	"encoding/json"
	"net/http"
	"sprint3/Helpers"
	"sprint3/Models"

	"github.com/gorilla/mux"
)

// CompanyAndOffice struct
type CompanyAndOffice struct {
	SaasCompanyID uint
	SaasOfficeID  uint
}

// GetSaasCompanyAndOffice handler
// @Summary retrieve a saascompany and saasOffice by driverID
// @Description retrieve a saascompany and saasOffice by driverID in database with the ID in request parameter
// @Tags saas_company_to_driver_relations saas_office_to_driver_relation
// @Accept  json
// @Produce  json
// @Param driver_id path integer true "enter driver_id:"
// @Success 200 {object} CompanyAndOffice
// @Router /saascompanyAndoffice/{driver_id} [get]
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

func GetSaasCompany(w http.ResponseWriter, r *http.Request) {

	db := Helpers.DbConnect()
	var saasCompany []Models.SaasCompanies
	db.Model(&saasCompany).Find(&saasCompany)
	json.NewEncoder(w).Encode(saasCompany)

}
