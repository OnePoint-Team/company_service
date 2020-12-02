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
	err := companyObject.Select(pathvar.ID)

	if err == nil {
		data := schemas.CompanySerializer(&companyObject)
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{"message": "not found"})
	}
}

// GetCompanies fetches all companies from database
func GetCompanies(c *gin.Context) {
	companyObject := company.Company{}
	var companies []company.Company

	err := companyObject.SelectAll(&companies)

	if err == nil {
		data := schemas.CompaniesSerializer(&companies)
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{"message": "not found"})
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
		c.JSON(422, gin.H{"message": "user not created"})
		return
	}

	//companyObject := company.Company{Name: obj["name"].(string)}
	company := company.Company{Name: input.Name}
	err := company.Insert()
	log.Println(company)

	if err != nil {
		c.JSON(404, gin.H{"message": "Failed to create"})
	} else {
		data := schemas.CompanySerializer(&company)
		c.JSON(200, data)
	}

}
