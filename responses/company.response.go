package responses

import "personal_hr/models"

// CompanyResponse ...
type CompanyResponse struct {
	models.BaseModels
	CompanyName string `json:"company_name"`
	models.BaseCUModels
}

// NewCompanyResponse ...
func NewCompanyResponse(m *models.Company) *CompanyResponse {
	return &CompanyResponse{
		BaseModels:   m.BaseModels,
		CompanyName:  m.CompanyName,
		BaseCUModels: m.BaseCUModels,
	}
}
