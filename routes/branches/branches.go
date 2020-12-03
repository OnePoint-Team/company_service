package branches

import (
	"log"
	"net/http"

	"github.com/OnePoint-Team/company_service/models/branch"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
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

	b := branch.Branch{Name: input.Name}
	if err := b.Insert(pathvar.ID); err != nil {
		log.Println("Insert error : ->", err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	data := schemas.BranchSerializer(&b)
	c.JSON(http.StatusOK, data)
	log.Println(b)

}

// GetBranches fetch all branches
func GetBranches(c *gin.Context) {
	var b branch.Branch
	listOfBranch := []branch.Branch{}
	var pathvar schemas.CompanyPathVar

	if err := c.BindUri(&pathvar); err != nil {
		log.Println(err)
		return
	}

	b.All(&listOfBranch, pathvar.ID)

	data := schemas.SerializeAllBranches(&listOfBranch)
	c.JSON(http.StatusOK, data)
}

// GetBranchByID fetch branch by id
func GetBranchByID(c *gin.Context) {
	var b branch.Branch
	var pathVar schemas.BranchPathVar

	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Branch id ->", pathVar.BID)
	log.Println("Company id ->", pathVar.CID)

	if err := b.Select(pathVar.BID, pathVar.CID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	data := schemas.BranchSerializer(&b)
	c.JSON(http.StatusOK, data)
}

// DeleteBranch remove branhc from db
func DeleteBranch(c *gin.Context) {
	var b branch.Branch
	var pathVar schemas.BranchPathVar

	if err := c.BindUri(&pathVar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "validation error"})
		log.Println(err.Error())
		return
	}
	// b.Select(pathVar.BID, pathVar.CID)
	b.Delete(pathVar.BID, pathVar.CID)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
