package schemas

import (
	"github.com/OnePoint-Team/company_service/models/branch"
)

// BranchCreate for input data
type BranchCreate struct {
	Name string `json:"name" binding:"required"`
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
