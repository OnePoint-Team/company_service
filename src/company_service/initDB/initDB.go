package initDB

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initDB function initialiazies db
func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=company_service port=5432 timezone=Asia/Baku"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("err ---> ", err)
	}

	return db
}
