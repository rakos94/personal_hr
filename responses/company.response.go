package responses

import (
	"personal_hr/helper"
	"personal_hr/models"
)

// CompanyResponse ...
type CompanyResponse struct {
	Departments []*DepartmentResponse `json:"-"`
	models.BaseModels
	CompanyName string `json:"company_name"`
	models.BaseCUModels
}

// NewCompanyResponse ...
func NewCompanyResponse(m *models.Company) *CompanyResponse {
	var departments []*DepartmentResponse
	for _, d := range m.Department {
		departments = append(departments, NewDepartmentResponse(&d))
	}
	return &CompanyResponse{
		Departments:  departments,
		BaseModels:   m.BaseModels,
		CompanyName:  m.CompanyName,
		BaseCUModels: m.BaseCUModels,
	}
}

// LoadRelation ...
func (m *CompanyResponse) LoadRelation() map[string]interface{} {
	res := helper.ConvertMap(m)
	res["departments"] = m.Departments
	return res
}
