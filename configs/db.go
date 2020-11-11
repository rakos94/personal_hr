package configs

import (
	"fmt"
	"personal_hr/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tables = []interface{}{
	&models.Person{},
	&models.Employee{},
	&models.Company{},
	&models.Department{},
	&models.Division{},
	&models.Location{},
	&models.Position{},
	&models.Assignement{},
	&models.City{},
	&models.Contry{},
	&models.Provices{},
}

func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, pass, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
