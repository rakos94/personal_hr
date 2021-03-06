package responses

import "personal_hr/models"

// DepartmentResponse ...
type DepartmentResponse struct {
	models.BaseModels
	DepartmentName string `json:"department_name"`
	CompanyID      string `json:"company_id"`
	models.BaseCUModels
}

// NewDepartmentResponse ...
func NewDepartmentResponse(m *models.Department) *DepartmentResponse {
	return &DepartmentResponse{
		BaseModels:     m.BaseModels,
		DepartmentName: m.DepartmentName,
		CompanyID:      m.CompanyID,
		BaseCUModels:   m.BaseCUModels,
	}
}

// NewListDepartmentResponse ...
func NewListDepartmentResponse(list []models.Department) []DepartmentResponse {
	var departments []DepartmentResponse
	for _, d := range list {
		departments = append(departments, *NewDepartmentResponse(&d))
	}
	return departments
}
