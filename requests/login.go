package requests

// Login ...
type Login struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `json:"-"`
	Token    string `json:"token" gorm:"-"`
}
