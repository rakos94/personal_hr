package requests

// PersonRequest ...
type PersonRequest struct {
	FirstName string  `form:"first_name" validate:"required"`
	LastName  *string `form:"last_name"`
	Email     string  `form:"email" validate:"required,email"`
	Password  string  `form:"password" validate:"required"`
	Gender    *string `form:"gender"`
}
