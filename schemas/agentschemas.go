package schemas

import "github.com/OnePoint-Team/company_service/models/agent"

// AgentCreate for body data
type AgentCreate struct {
	UserID string `json:"user_id" binding:"required,uuid4"`
}

// AgentPathVar struct
type AgentPathVar struct {
	BID string `uri:"bid" binding:"required,uuid4"`
	CID string `uri:"cid" binding:"required,uuid4"`
	AID string `uri:"aid" binding:"required,uuid4"`
}

// AgentSerializer serialize agent struct
func AgentSerializer(a agent.Agent) map[string]interface{} {
	data := make(map[string]interface{})
	data["id"] = a.Base.ID
	data["company_id"] = a.CompanyID
	data["branch_id"] = a.BranchID
	data["user_id"] = a.UserID
	data["created"] = a.Base.CreatedAt
	data["updated"] = a.Base.UpdatedAt
	return data
}

// SerializeAllAgents func
func SerializeAllAgents(b []agent.Agent) []map[string]interface{} {
	data := []map[string]interface{}{}
	for _, v := range b {
		temp := make(map[string]interface{})
		temp["id"] = v.Base.ID
		temp["created"] = v.Base.CreatedAt
		temp["updated"] = v.Base.UpdatedAt
		temp["branch_id"] = v.BranchID
		temp["company_id"] = v.CompanyID
		temp["user_id"] = v.UserID
		data = append(data, temp)
	}
	return data
}
