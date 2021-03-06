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
	&models.Country{},
	&models.Provinces{},
}

func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, pass, dbname, port, sslmode)
	//postgresql://103.30.180.34:9595/booting#spring.datasource.username=postgres#spring.datasource.password=bootlawen123
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
