package requests

import (
	"database/sql"
)

// PersonRequest ...
type PersonRequest struct {
	FirstName string         `form:"first_name" validate:"required"`
	LastName  sql.NullString `form:"last_name"`
	Email     string         `form:"email" validate:"required,email"`
	Password  string         `form:"password" validate:"required"`
	Gender    sql.NullString `form:"gender"`
}
