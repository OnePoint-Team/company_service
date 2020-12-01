package main

import (
	"github.com/OnePoint-Team/company_service/configs"
	"github.com/OnePoint-Team/company_service/routes/companies"
	"github.com/gin-gonic/gin"
)

func main() {
	// migrate()

	r := gin.Default()
	r.GET("/companies", companies.GetCompanies)
	r.POST("/companies", companies.POSTCompanies)

	r.GET("/companies/:id", companies.GetByID)

	r.Run(configs.Config.Host + ":" + configs.Config.Port)

}
