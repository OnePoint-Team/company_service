package main

import (
	"github.com/OnePoint-Team/company_service/configs"
	"github.com/OnePoint-Team/company_service/routes/agents"
	"github.com/OnePoint-Team/company_service/routes/branches"
	"github.com/OnePoint-Team/company_service/routes/companies"
	"github.com/OnePoint-Team/company_service/routes/lenders"
	"github.com/gin-gonic/gin"
)

// MappingUrls map url to group
func MappingUrls() *gin.Engine {
	router := gin.Default()
	r := router.Group("/companies")
	{
		r.GET("/", companies.GetCompanies)
		r.POST("/", companies.CreateCompanies)
		r.GET("/:cid", companies.GetByID)
		r.POST("/:cid/branches", branches.CreateBranch)
		r.GET("/:cid/branches", branches.GetBranches)
		r.GET("/:cid/branches/:bid", branches.GetBranchByID)
		r.DELETE("/:cid/branches/:bid", branches.DeleteBranch)
		r.POST("/:cid/branches/:bid/agents", agents.CreateAgent)
		r.GET("/:cid/branches/:bid/agents", agents.GetAgents)
		r.GET("/:cid/branches/:bid/agents/:aid", agents.GetAgentByID)
		r.PUT("/:cid/branches/:bid/agents/:aid", agents.UpdateAgent)
		r.DELETE("/:cid/branches/:bid/agents/:aid", agents.DeleteAgent)
	}
	l := router.Group("/lenders")
	{
		l.POST("/", lenders.CreateLender)
		l.GET("/:lid", lenders.GetLender)
	}
	return router
}

func main() {
	// migrations.Migrate()

	r := MappingUrls()

	r.Run(configs.Config.Host + ":" + configs.Config.Port)

}
