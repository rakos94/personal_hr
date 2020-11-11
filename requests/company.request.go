package requests

// CompanyRequest ...
type CompanyRequest struct {
	CompanyName string `form:"company_name" validate:"required"`
}
