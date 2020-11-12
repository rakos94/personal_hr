package requests

import (
	"gorm.io/datatypes"
)

// PersonRequest ...
type PersonRequest struct {
	FirstName         string         `form:"first_name" validate:"required"`
	LastName          *string        `form:"last_name"`
	Email             string         `form:"email" validate:"required,email"`
	Password          string         `form:"password" validate:"required"`
	BirthPlace        *string        `form:"birth_place"`
	BirthDate         datatypes.Date `form:"birth_date" validate:"required"`
	Mobile            *string        `form:"mobile"`
	Gender            string         `form:"gender" validate:"required"`
	IsFromRecruitment bool           `form:"is_from_recruitment" validate:"required"`
	CreatedBy         string         `form:"created_by" validate:"required"`
	UpdatedBy         string         `form:"updated_by" validate:"required"`
}
