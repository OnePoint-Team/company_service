package companies

import (
	"github.com/OnePoint-Team/company_service/models/company"
	"github.com/OnePoint-Team/company_service/schemas"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

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

// POSTCompanies gets
func POSTCompanies(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var obj map[string]interface{}
	err := json.Unmarshal(body, &obj)

	if err != nil {
		log.Println("Error occuried")
		c.JSON(404, gin.H{"message": "user not created"})
	}

	companyObject := company.Company{Name: obj["name"].(string)}
	result := companyObject.Insert()
	log.Println(companyObject)

	if result.Error != nil {
		c.JSON(404, gin.H{"message": "Failed to create"})
	} else {
		data := schemas.CompanySerializer(&companyObject)
		c.JSON(200, data)
	}

}
