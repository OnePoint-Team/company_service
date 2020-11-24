package companyschemas

import (
	"company_service/models/company"
	"encoding/json"
	"fmt"
	"log"
)

// type serializedCompany struct {
// 	ID        uuid.UUID          `json:"id"`
// 	Name      string             `json:"name"`
// 	CreatedAt time.Time          `json:"created"`
// 	UpdatedAt time.Time          `json:"updated"`
// 	Branches  []serializedBranch `json:"branches"`
// }

// type serializedBranch struct {
// 	ID        uuid.UUID `json:"id"`
// 	CreatedAt time.Time `json:"created"`
// 	UpdatedAt time.Time `json:"updated"`
// 	Name      string    `json:"name"`
// 	CompanyID uuid.UUID `json:"company_id"`
// }

// Serializer to serialize object
func Serializer(c *company.Company) {
	var jsonData []byte
	jsonData, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("jsonData --->", string(jsonData))

}