package companies

import (
	"log"

	"github.com/OnePoint-Team/company_service/models/company"
	"github.com/OnePoint-Team/company_service/schemas"

	// uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"

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
	validate := validator.New()
	if err := validate.Struct(pathvar); err != nil {
		c.SecureJSON(400, gin.H{"message": "not found"})
		return
	}
	result := companyObject.Select(pathvar.ID)

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
	// body, _ := ioutil.ReadAll(c.Request.Body)
	var input schemas.CompanyCreate
	// var obj map[string]interface{}
	// err := json.Unmarshal(body, &obj)

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Error occuried")
		c.JSON(422, gin.H{"message": "user not created"})
		return
	}

	//companyObject := company.Company{Name: obj["name"].(string)}
	company := company.Company{Name: input.Name}
	err := company.Insert()
	log.Println(company)

	if err.Error != nil {
		c.JSON(404, gin.H{"message": "Failed to create"})
	} else {
		data := schemas.CompanySerializer(&company)
		c.JSON(200, data)
	}

}
