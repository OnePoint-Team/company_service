package schemas

import (
	"company_service/models/company"
)

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
