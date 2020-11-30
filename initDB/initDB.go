package initDB

import (
	"fmt"

	"github.com/OnePoint-Team/company_service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB function initialiazies db
func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.Config.DSN), &gorm.Config{})

	if err != nil {
		fmt.Println("err ---> ", err)
	}

	return db
}
