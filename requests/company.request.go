package requests

// CompanyRequest ...
type CompanyRequest struct {
	CompanyName string `json:"company_name" validate:"required"`
}
