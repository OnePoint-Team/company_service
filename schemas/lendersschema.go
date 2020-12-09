package schemas

import "github.com/OnePoint-Team/company_service/models/lender"

// LenderCreate body data for lender create
type LenderCreate struct {
	Name string `json:"name" binding:"required"`
}

// LenderPath for lenders path variable
type LenderPath struct {
	LID string `uri:"lid" binding:"required,uuid4"`
}

// LenderSerializer for lender model
func LenderSerializer(l lender.Lender) map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = l.Base.ID
	m["name"] = l.Name
	m["created"] = l.Base.CreatedAt
	m["updated"] = l.Base.UpdatedAt

	return m
}
