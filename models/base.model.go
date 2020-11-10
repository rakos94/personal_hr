package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModels struct {
	Id string `json:"id" gorm:"primaryKey"`
}

type BaseCUModels struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (base *BaseModels) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	base.Id = uuid.String()
	return nil
}
