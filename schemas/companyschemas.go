package schemas

import (
	"github.com/OnePoint-Team/company_service/models/company"
)

// CompanyCreate for input data
type CompanyCreate struct {
	Name string `json:"name" binding:"required"`
}

//CompanyPathVar for path variable
type CompanyPathVar struct {
	ID string `uri:"id" binding:"required" validate:"required,uuid4"`
}

// CompanySerializer to serialize object
func CompanySerializer(c *company.Company) map[string]interface{} {
	data := make(map[string]interface{})

	data["name"] = c.Name
	data["id"] = c.Base.ID.String()
	data["created"] = c.Base.CreatedAt.String()
	data["updated"] = c.Base.UpdatedAt.String()

	return data
}

// CompaniesSerializer to serialize object
func CompaniesSerializer(c *[]company.Company) map[string][]interface{} {
	data := make(map[string][]interface{})

	for _, value := range *c {
		company := CompanySerializer(&value)
		data["companies"] = append(data["companies"], company)

		for _, value := range value.Branches {
			data["branches"] = append(data["branches"], BranchSerializer(&value))
		}

	}
	return data
}
