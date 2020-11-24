package main

import (
	"company_service/initDB"
	"company_service/models/branch"
	"company_service/models/company"
	"company_service/routes/companies"

	"github.com/gin-gonic/gin"
)

func migrate() {
	// Migration
	db := initDB.InitDB()
	db.AutoMigrate(&company.Company{}, &branch.Branch{})

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
	// c := company.Company{Name: "Optimal"}
	// c.Insert()

	// d := company.Company{Name: "BakuElectronics"}
	// d.Insert()

	// z := company.Company{Name: "Kontakt"}
	// z.Insert()

	r := gin.Default()
	r.GET("/", companies.GetCompanies)
	r.GET("/:id", companies.GetByID)
	r.Run()

}
