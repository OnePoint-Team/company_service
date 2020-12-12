package operators

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OnePoint-Team/company_service/models/operators"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetOperators(c *gin.Context) {
	operatorObject := operators.Operators{}

	var operators []operators.Operators

	err := operatorObject.All(&operators)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	}
	c.JSON(http.StatusOK, operators)

}

func CreateOperator(c *gin.Context) {
	var input schemas.OperatorCreate
	var pathVar schemas.OperatorPathVar

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		return
	}
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}

	operator := operators.Operators{}

	err := operator.Insert(pathVar.LID, input.UserID)
	log.Println(operator)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to create"})

		return
	}

	c.JSON(http.StatusNotFound, operator)
}

func GetOneOperator(c *gin.Context) {
	var pathVar schemas.OperatorPath

	if err := c.BindUri(&pathVar); err != nil {
		log.Println("ERRRROOR", err)
		return
	}
	fmt.Println(pathVar)
	operatorObject := operators.Operators{}

	err := operatorObject.Select(pathVar.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, operatorObject)

}
func DeleteOperator(c *gin.Context) {
	var pathVar schemas.OperatorPath

	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
	}

	operatorObject := operators.Operators{}

	err := operatorObject.Delete(pathVar.UserID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Deleted"})

}

func UpdateOperator(c *gin.Context) {

	var pathVar schemas.OperatorPath
	var bodyVar schemas.OperatorCreate
	
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}

	if err := c.Bind(&bodyVar); err != nil {
		log.Println(err)
		return
	}
	 operator  := operators.Operators{}

	operator.UserID, _ = uuid.FromString(bodyVar.UserID)
	if err := operator.Update(pathVar.UserID); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"result": true})

}
