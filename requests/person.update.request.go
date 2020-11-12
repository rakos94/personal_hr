package requests

import "time"

// PersonUpdateRequest ...
type PersonUpdateRequest struct {
	FirstName         string    `form:"first_name"`
	LastName          *string   `form:"last_name"`
	Email             string    `form:"email" validate:"email"`
	BirthPlace        *string   `form:"birth_place"`
	BirthDate         time.Time `form:"birth_date"`
	Mobile            *string   `form:"mobile"`
	Gender            string    `form:"gender"`
	IsFromRecruitment bool      `form:"is_from_recruitment"`
	CreatedBy         string    `form:"created_by"`
	UpdatedBy         string    `form:"updated_by"`
}
