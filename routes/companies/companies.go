package companies

import (
	"log"
	"net/http"

	"github.com/OnePoint-Team/company_service/models/company"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
)

// GetByID fetches company by id from database
func GetByID(c *gin.Context) {
	companyObject := company.Company{}
	var pathvar schemas.CompanyPathVar

	// id := c.Param("id")
	if err := c.BindUri(&pathvar); err != nil {
		log.Println(err)
		return
	}
	err := companyObject.Select(pathvar.ID)

	if err == nil {
		// ###################
		// serializer.Schema(companyObject)

		// data := schemas.CompanySerializer(&companyObject)

		// serializer.Schema(companyObject)
		c.JSON(http.StatusOK, companyObject)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
}

// GetCompanies fetches all companies from database
func GetCompanies(c *gin.Context) {
	companyObject := company.Company{}
	var companies []company.Company

	err := companyObject.All(&companies)

	if err == nil {
		// data := []map[string]interface{}{}
		// data := schemas.CompaniesSerializer(&companies)

		c.JSON(http.StatusOK, companies)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
}

// CreateCompanies gets
func CreateCompanies(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	var input schemas.CompanyCreate
	// var obj map[string]interface{}
	// err := json.Unmarshal(body, &obj)

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		return
	}

	//companyObject := company.Company{Name: obj["name"].(string)}
	company := company.Company{Name: input.Name}
	err := company.Insert()
	log.Println(company)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to create"})
	} else {
		// data := schemas.CompanySerializer(&company)
		c.JSON(http.StatusOK, company)
	}

}
