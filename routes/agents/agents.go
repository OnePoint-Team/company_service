package agents

import (
	"log"
	"net/http"

	"github.com/OnePoint-Team/company_service/models/agent"
	"github.com/OnePoint-Team/company_service/schemas"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetAgents fetch all agents
func GetAgents(c *gin.Context) {
	var a agent.Agent
	var agents []agent.Agent
	var pathVar schemas.BranchPathVar
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}

	a.All(&agents, pathVar.CID, pathVar.BID)
	data := schemas.SerializeAllAgents(agents)
	c.JSON(http.StatusOK, data)
}

// GetAgentByID get agent
func GetAgentByID(c *gin.Context) {
	var a agent.Agent
	var pathVar schemas.AgentPathVar
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}
	if err := a.Select(pathVar.AID, pathVar.CID, pathVar.BID); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail"})
		return
	}
	data := schemas.AgentSerializer(a)
	c.JSON(http.StatusOK, data)
}

// CreateAgent func
func CreateAgent(c *gin.Context) {
	var a agent.Agent
	var pathVar schemas.BranchPathVar
	var bodyVar schemas.AgentCreate
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}
	if err := c.Bind(&bodyVar); err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v %+v", pathVar, bodyVar)
	if err := a.Insert(pathVar.CID, pathVar.BID, bodyVar.UserID); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data := schemas.AgentSerializer(a)
	c.JSON(http.StatusOK, data)
}

// UpdateAgent update
func UpdateAgent(c *gin.Context) {
	var a agent.Agent
	var pathVar schemas.AgentPathVar
	var bodyVar schemas.AgentCreate
	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}
	if err := c.Bind(&bodyVar); err != nil {
		log.Println(err)
		return
	}
	if err := a.Select(pathVar.AID, pathVar.CID, pathVar.BID); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail"})
		return
	}
	a.UserID, _ = uuid.FromString(bodyVar.UserID)
	if err := a.Update(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail"})
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}

// DeleteAgent delete
func DeleteAgent(c *gin.Context) {
	var a agent.Agent
	var pathVar schemas.AgentPathVar

	if err := c.BindUri(&pathVar); err != nil {
		log.Println(err)
		return
	}

	if err := a.Select(pathVar.AID, pathVar.CID, pathVar.BID); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail"})
		return
	}

	if err := a.Delete(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail"})
	}
	c.JSON(http.StatusOK, gin.H{"result": true})
}
