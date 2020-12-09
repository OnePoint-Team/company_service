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

func splitter(s, delimeter string) []string {
	arr := strings.Split(s, delimeter)
	return arr
}

// Schema function
func Schema(i interface{}) {
	outPutSchema := make(map[string]interface{})
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	n := t.NumField()

	for i := 0; i < n; i++ {
		tt := t.Field(i)

		tagName := tt.Tag.Get("responseSchema")
		if tagName != "" {
			switch tt.Type.Kind() {
			case reflect.Struct:

				if strings.Contains(tagName, ",") {
					for _, str := range strings.Split(tagName, ",") {
						value := strings.Split(str, ":")
						outPutSchema[value[0]] = v.Field(i).FieldByName(value[1]).Interface()
					}
				}

			case reflect.String:
				val := v.Field(i)
				outPutSchema[tagName] = val
			}

		}
	}
	fmt.Println(outPutSchema)
}

// tagName:  id:Base.ID,created:Base.created
// v.Field(i) :  {f867fa47-ed88-44c0-9635-ff2ebf176231 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
