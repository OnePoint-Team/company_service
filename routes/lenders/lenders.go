package lenders

import (
	"log"
	"net/http"

	"github.com/OnePoint-Team/company_service/models/lender"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
)

func CreateLender(c *gin.Context) {
	var bodyData schemas.LenderCreate
	if err := c.ShouldBindJSON(bodyData); err != nil {
		log.Println("error body data")
		return
	}
	l := lender.Lender{Name: bodyData.Name}
	if err := l.Insert(); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Fail to create"})
		return
	}
}
