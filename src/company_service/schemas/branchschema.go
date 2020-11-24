package schemas

import (
	"company_service/models/branch"
	"encoding/json"
	"fmt"
	"log"
)

// Serializer to serialize object
func BranchSerializer(b *branch.Branch) {
	var jsonData []byte
	jsonData, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	for key, value := range jsonData {
		fmt.Println("Key:", key, "Value:", value)
	}
	// fmt.Println("jsonData --->", string(jsonData))

}
