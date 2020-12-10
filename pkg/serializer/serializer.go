package serializer

import (
	"encoding/json"
	"fmt"
)

// func convertStructToMap(i interface{}) map[string]interface{} {
// 	var m map[string]interface{}
// 	b, _ := json.Marshal(i)
// 	json.Unmarshal([]byte(string(b)), &m)
// 	return m
// }

// func splitter(s, delimeter string) []string {
// 	arr := strings.Split(s, delimeter)
// 	return arr
// }

// // Schema function
// func Schema(i interface{}) map[string]interface{} {
// 	outPutSchema := make(map[string]interface{})
// 	t := reflect.TypeOf(i)
// 	v := reflect.ValueOf(i)
// 	n := t.NumField()
// 	for i := 0; i < n; i++ {
// 		tt := t.Field(i)

// 		tagName := tt.Tag.Get("responseSchema")
// 		if tagName != "" {
// 			switch tt.Type.Kind() {
// 			case reflect.Struct:

// 				if strings.Contains(tagName, ",") {
// 					for _, str := range strings.Split(tagName, ",") {
// 						value := strings.Split(str, ":")
// 						outPutSchema[value[0]] = v.Field(i).FieldByName(value[1]).Interface()
// 					}
// 				}
// 			case reflect.Slice:
// 				// SchemaSlice(v.Field(i).Interface())
// 				if v.Field(i).Len() != 0 {
// 					// fmt.Println("V -> ", v.Field(i).Index(0).FieldByName("Name").Interface())
// 					temp := []interface{}{}
// 					if strings.Contains(tagName, ",") {
// 						for _, str := range strings.Split(tagName, ",") {
// 							value := strings.Split(str, ":")
// 							fmt.Println("Value -> ", value)
// 							for j := 0; j < v.Field(i).Len(); j++ {
// 								temp = append(temp, v.Field(i).Index(0).FieldByName(value[1]).Interface())
// 							}
// 							outPutSchema[value[0]] = temp
// 						}
// 					} else {
// 						value := strings.Split(tagName, ":")
// 						for j := 0; j < v.Field(i).Len(); j++ {

// 							temp = append(temp, v.Field(i).Index(j).FieldByName(value[1]).Interface())

// 						}
// 						outPutSchema[value[0]] = temp
// 					}
// 				}
// 				// fmt.Println("TEMP -> ", t.Field(i).Field(i).Interface())
// 			case reflect.String:
// 				val := v.Field(i).String()
// 				outPutSchema[tagName] = val
// 			default:
// 				fmt.Println("Unknown types -> ", tt)
// 			}

// 		}
// 	}
// 	fmt.Println("OUTPUT Schema -> ", outPutSchema)
// 	return outPutSchema
// }

// // SchemaSlice func
// func SchemaSlice(i interface{}) []map[string]interface{} {
// 	outPutSchema := []map[string]interface{}{}
// 	v := reflect.ValueOf(i)
// 	t := reflect.TypeOf(i)
// 	switch t.Kind() {
// 	case reflect.Slice:

// 		for i := 0; i < v.Len(); i++ {
// 			temp := Schema(v.Index(i).Interface())
// 			outPutSchema = append(outPutSchema, temp)
// 		}
// 	}
// 	return outPutSchema
// }

// // tagName:  id:Base.ID,created:Base.created
// // v.Field(i) :  {f867fa47-ed88-44c0-9635-ff2ebf176231 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}

type Base struct {
	ID      string
	created string
	updated string
}

// CompanySchema struct
type CompanySchema struct {
	name string
	Base
}

type BranchSchema struct {
	Base Base
}

//CompaniesSchema struct
type CompaniesSchema struct {
	name     string
	id       string
	created  string
	updated  string
	branches []BranchSchema
}

func Schema(i interface{}) {
	// js, _ := json.Marshal(i)
	x, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(x))
	return
}
