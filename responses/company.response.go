package responses

import (
	"personal_hr/helper"
	"personal_hr/models"
)

// CompanyResponse ...
type CompanyResponse struct {
	Departments []DepartmentResponse `json:"-"`
	models.BaseModels
	CompanyName string `json:"company_name"`
	models.BaseCUModels
}

// NewCompanyResponse ...
func NewCompanyResponse(m *models.Company) *CompanyResponse {
	return &CompanyResponse{
		Departments:  NewListDepartmentResponse(m.Department),
		BaseModels:   m.BaseModels,
		CompanyName:  m.CompanyName,
		BaseCUModels: m.BaseCUModels,
	}
}

// NewListCompanyResponse ...
func NewListCompanyResponse(list []models.Company) []CompanyResponse {
	var companies []CompanyResponse
	for _, d := range list {
		companies = append(companies, *NewCompanyResponse(&d))
	}
	return companies
}

// LoadRelation ...
func (m *CompanyResponse) LoadRelation() map[string]interface{} {
	res := helper.ConvertMap(m)
	res["departments"] = m.Departments
	return res
}
