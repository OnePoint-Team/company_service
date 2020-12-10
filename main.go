package main

import (
	"github.com/OnePoint-Team/company_service/configs"
	"github.com/OnePoint-Team/company_service/docs"
	"github.com/OnePoint-Team/company_service/initdb"
	"github.com/OnePoint-Team/company_service/initdb/migrations"
	"github.com/OnePoint-Team/company_service/routes/agents"
	"github.com/OnePoint-Team/company_service/routes/branches"
	"github.com/OnePoint-Team/company_service/routes/companies"
	"github.com/OnePoint-Team/company_service/routes/lenders"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// "github.com/OnePoint-Team/company_service/middleware"
)

// MappingUrls map url to group
func MappingUrls() *gin.Engine {

	//in case if we want to achive  middlare and some customization we do followings NOTE;
	// router := gin.New()
	// router.Use(gin.Recovery(), middleware.Logger())

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
		l.GET("/", lenders.GetAllLenders)
		l.GET("/:lid", lenders.GetLender)
		l.DELETE("/:lid", lenders.Delete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func main() {

	docs.SwaggerInfo.Title = "Company Service "
	docs.SwaggerInfo.Description = "Company Service "
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer initdb.CloseDB(initdb.DbInstance)

	migrations.Migrate()

	r := MappingUrls()

	r.Run(configs.Config.Host + ":" + configs.Config.Port)

}
