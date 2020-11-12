package responses

import "personal_hr/models"

// DivisionResponse ...
type DivisionResponse struct {
	models.BaseModels
	DivisionName string `json:"division_name"`
	DepartmentID string `json:"department_id"`
	models.BaseCUModels
}

// NewDivisionResponse ...
func NewDivisionResponse(m *models.Division) *DivisionResponse {
	return &DivisionResponse{
		BaseModels:   m.BaseModels,
		DivisionName: m.DivisionName,
		DepartmentID: m.DepartmentID,
		BaseCUModels: m.BaseCUModels,
	}
}
