package main

import (
	"company_service/initDB"
	"company_service/models/branch"
	"company_service/models/company"
	"fmt"
)

func migrate() {
	// Migration
	db := initDB.InitDB()
	err := db.AutoMigrate(&company.Company{}, &branch.Branch{})
	fmt.Println("err->", err)

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
	fmt.Fprintln("hes")
	// r := gin.Default()
	// r.GET("/", companies.GetCompanies)
	// r.POST("/", companies.POSTCompanies)

	// r.GET("/:id", companies.GetByID)

	// r.Run()

}
