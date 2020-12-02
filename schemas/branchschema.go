package schemas

import (
	"github.com/OnePoint-Team/company_service/models/branch"
)

// BranchCreate for input data
type BranchCreate struct {
	Name string `json:"name" binding:"required"`
}

// BranchPathVar for path variable
type BranchPathVar struct {
	BID string `uri:"bid" binding:"required" validate:"required,uuid4"`
	CID string `uri:"cid" binding:"required" validate:"required,uuid4"`
}

// BranchSerializer to serialize Branch object
func BranchSerializer(b *branch.Branch) map[string]interface{} {
	data := make(map[string]interface{})

	data["id"] = b.Base.ID
	data["created"] = b.Base.CreatedAt
	data["updated"] = b.Base.UpdatedAt
	data["company_id"] = b.CompanyID
	data["name"] = b.Name
	return data
}

// SerializeAllBranches test
func SerializeAllBranches(b *[]branch.Branch) []map[string]interface{} {
	data := []map[string]interface{}{}
	for _, v := range *b {
		temp := make(map[string]interface{})
		temp["id"] = v.Base.ID
		temp["created"] = v.Base.CreatedAt
		temp["updated"] = v.Base.UpdatedAt
		temp["name"] = v.Name
		temp["company_id"] = v.CompanyID
		data = append(data, temp)
	}
	return data
}
