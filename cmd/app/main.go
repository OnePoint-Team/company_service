package main

import (
	"github.com/OnePoint-Team/company_service/configs"
	"github.com/OnePoint-Team/company_service/routes/branches"
	"github.com/OnePoint-Team/company_service/routes/companies"
	"github.com/gin-gonic/gin"
)

func main() {
	// migrations.Migrate()

	r := gin.Default()
	r.GET("/companies", companies.GetCompanies)
	r.POST("/companies", companies.CreateCompanies)

	r.GET("/companies/:cid", companies.GetByID)
	r.POST("/companies/:cid/branches", branches.CreateBranch)
	r.GET("/companies/:cid/branches", branches.GetBranches)
	r.GET("/companies/:cid/branches/:bid", branches.GetBranchByID)

	r.Run(configs.Config.Host + ":" + configs.Config.Port)

}
