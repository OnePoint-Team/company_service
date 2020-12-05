package serializer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func convertStructToMap(i interface{}) map[string]interface{} {
	var m map[string]interface{}
	b, _ := json.Marshal(i)
	json.Unmarshal([]byte(string(b)), &m)
	return m
}

func splitter(s string) (string, string) {
	arr := strings.Split(s, ":")
	return arr[0], arr[1]
}

func outputSchema(jsonData, outputData map[string]interface{}, tagName string) {
	tags := strings.Split(tagName, ",")

	for _, v := range tags {
		fmt.Println(v)
		if strings.Contains(v, ":") {
			key, val := splitter(v)
			t := strings.Split(val, ".")
			u := jsonData[t[0]]
			y := u.(map[string]interface{})

			outputData[key] = y[t[1]]

			// outputData[key] =
		}
	}
}

// Schema function
func Schema(i interface{}) {
	t := reflect.TypeOf(i)
	jsonData := convertStructToMap(i)
	fmt.Println(jsonData)
	outputData := make(map[string]interface{})

	n := t.NumField()

	for i := 0; i < n; i++ {
		tt := t.Field(i)
		tagName := tt.Tag.Get("serialize")
		if tagName != "" {
			outputSchema(jsonData, outputData, tagName)
		}
	}
	fmt.Println("output -> ", outputData)
}
