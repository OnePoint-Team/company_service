package branches

import (
	"log"

	"github.com/OnePoint-Team/company_service/models/branch"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// CreateBranch for creation branches
func CreateBranch(c *gin.Context) {
	var input schemas.BranchCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		return
	}
	var pathvar schemas.CompanyPathVar

	if err := c.BindUri(&pathvar); err != nil {
		log.Println(err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(pathvar); err != nil {
		c.SecureJSON(400, gin.H{"message": "not found"})
		return
	}

	b := branch.Branch{Name: input.Name}
	if err := b.Insert(pathvar.ID); err != nil {
		log.Println("Insert error : ->", err)
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	data := schemas.BranchSerializer(&b)
	c.JSON(200, data)
	log.Println(b)

}
