package companies

import (
	"company_service/models/company"
	"company_service/schemas/companyschemas"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetCompanies fetches all companies from database
func GetCompanies(c *gin.Context) {
	companyObject := company.Company{}
	var companies []company.Company

	result := companyObject.SelectAll(&companies)
	if result.Error == nil {
		c.JSON(200, gin.H{
			"companies": companies,
		})
	} else {
		c.JSON(404, gin.H{"message": "not found"})
	}
}

// GetByID fetches company by id from database
func GetByID(c *gin.Context) {
	companyObject := company.Company{}
	id := c.Param("id")
	result := companyObject.Select(id)

	companyschemas.Serializer(&companyObject)

	fmt.Println("HEREEEEE")
	if result.Error == nil {
		c.SecureJSON(200, companyObject)
	} else {
		c.SecureJSON(404, gin.H{"message": "not found"})
	}
}
