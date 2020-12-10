package initdb

import (
	"github.com/OnePoint-Team/company_service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbInstance connection instance
var DbInstance *gorm.DB

// InitDB function initialiazies db
func init() {
	var err error

	DbInstance, err = gorm.Open(postgres.Open(configs.Config.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func CloseDB(instance *gorm.DB) {

	db, err := instance.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	db.Close()
}
