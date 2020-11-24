package companies

import (
	"company_service/models/company"
	"company_service/schemas"

	"github.com/gin-gonic/gin"
)

// GetCompanies fetches all companies from database
func GetCompanies(c *gin.Context) {
	companyObject := company.Company{}
	var companies []company.Company

	result := companyObject.SelectAll(&companies)

	if result.Error == nil {
		schemas.CompaniesSerializer(&companies)
	}

	if result.Error == nil {
		data := schemas.CompaniesSerializer(&companies)
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{"message": "not found"})
	}
}

// GetByID fetches company by id from database
func GetByID(c *gin.Context) {
	companyObject := company.Company{}
	id := c.Param("id")
	result := companyObject.Select(id)

	if result.Error == nil {
		data := schemas.CompanySerializer(&companyObject)
		c.SecureJSON(200, data)
	} else {
		c.SecureJSON(404, gin.H{"message": "not found"})
	}
}
