package schemas

import (
	"github.com/OnePoint-Team/company_service/models/branch"
)

// BranchSerializer to serialize Branch object
func BranchSerializer(b *branch.Branch) map[string]interface{} {
	data := make(map[string]interface{})

	data["id"] = b.Base.ID
	data["created"] = b.Base.CreatedAt
	data["created"] = b.Base.UpdatedAt
	data["company_id"] = b.CompanyID
	data["name"] = b.Name
	return data
}
