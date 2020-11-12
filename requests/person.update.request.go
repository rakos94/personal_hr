package requests

// PersonUpdateRequest ...
type PersonUpdateRequest struct {
	FirstName string  `form:"first_name"`
	LastName  *string `form:"last_name"`
	Email     string  `form:"email" validate:"email"`
	Gender    *string `form:"gender"`
}
