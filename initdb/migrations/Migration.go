package migrations

import (
	"fmt"

	"github.com/OnePoint-Team/company_service/initdb"
	"github.com/OnePoint-Team/company_service/models/agent"
	"github.com/OnePoint-Team/company_service/models/branch"
	"github.com/OnePoint-Team/company_service/models/company"
	"github.com/OnePoint-Team/company_service/models/lender"
	"github.com/OnePoint-Team/company_service/models/operators"
)

// Migrate migrate db
func Migrate() {

	var models = []interface{}{&company.Company{}, &branch.Branch{}, &agent.Agent{},&lender.Lender{}, &operators.Operators{}}

	err := initdb.DbInstance.AutoMigrate(models...)

	if err != nil {
		panic(err)
	}
	fmt.Println("Migration Successfull->")

}

func InsertIntoCompany() {
	c := company.Company{Name: "_(*(*^*;"}
	c.Insert()
}

func InsertIntoBranch() {

	c := company.Company{}
	b := branch.Branch{Name: "Baku"}

	initdb.DbInstance.First(&c)
	b.CompanyID = c.Base.ID

	// b.Insert()
}

func selectTest() {

	c := company.Company{}

	initdb.DbInstance.Select("name; drop table users;").First(&c)
	// db.Select("name; drop table users;").First(&user)
}
