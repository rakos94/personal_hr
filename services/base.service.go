package services

import (
	"gorm.io/gorm"
)

var g *gorm.DB

func SetService(dg *gorm.DB) {
	g = dg
}
