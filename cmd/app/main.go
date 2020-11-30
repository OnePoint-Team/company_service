package main

import (
	"fmt"

	"github.com/OnePoint-Team/company_service/initDB"
	"github.com/OnePoint-Team/company_service/models/agent"
	"github.com/OnePoint-Team/company_service/models/branch"
	"github.com/OnePoint-Team/company_service/models/company"
)

func migrate() {
	// Migration
	db := initDB.InitDB()
	err := db.AutoMigrate(&company.Company{}, &branch.Branch{}, &agent.Agent{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Migration Successfull->")

}

func insert_into_company() {
	c := company.Company{Name: "_(*(*^*;"}
	c.Insert()
}

func insert_into_branch() {
	db := initDB.InitDB()

	c := company.Company{}
	b := branch.Branch{Name: "Baku"}

	db.First(&c)
	b.CompanyID = c.Base.ID

	b.Insert()
}

func select_test() {
	db := initDB.InitDB()
	c := company.Company{}

	db.Select("name; drop table users;").First(&c)
	// db.Select("name; drop table users;").First(&user)
}

func main() {
	migrate()

	// r := gin.Default()
	// r.GET("/", companies.GetCompanies)
	// r.POST("/", companies.POSTCompanies)

	// r.GET("/:id", companies.GetByID)

	// r.Run(configs.Config.Host + ":" + configs.Config.Port)

}
