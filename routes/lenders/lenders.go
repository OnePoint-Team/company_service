package lenders

import (
	"log"
	"net/http"

	"github.com/OnePoint-Team/company_service/models/lender"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
)

// CreateLender create lender
func CreateLender(c *gin.Context) {
	var bodyData schemas.LenderCreate
	if err := c.ShouldBind(&bodyData); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to create"})
		return
	}
	l := lender.Lender{Name: bodyData.Name}
	if err := l.Insert(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to create"})
		return
	}
	data := schemas.LenderSerializer(l)
	c.JSON(http.StatusOK, data)
}

// GetLender get lender by ID
func GetLender(c *gin.Context) {
	var pathVar schemas.LenderPath
	var lender lender.Lender
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}
	if err := lender.Select(pathVar.LID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	data := schemas.LenderSerializer(lender)
	c.JSON(http.StatusOK, data)

}
