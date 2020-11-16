package dao

import "personal_hr/models"

// AddressDao ...
type AddressDao interface {
	CreateAddress(data *models.Address) (*models.Address, error)
	GetAddressAll() ([]models.Address, error)
	GetAddressByID(id string) (models.Address, error)
	UpdateAddress(id string, data *models.Address) (*models.Address, error)
}
