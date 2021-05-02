package lichess

import (
	"reflect"
)

func fetchIds(v interface{}) string {
	var ids_str string
	ref_v := reflect.ValueOf(v)
	switch v.(type) {
	case []string:
		for i := 0; i < ref_v.Len(); i++ {
			ids_str += ref_v.Index(i).String()
			if i < ref_v.Len()-1 {
				ids_str += ","
			}
		}
	case []User:
		for i := 0; i < ref_v.Len(); i++ {
			ids_str += ref_v.Index(i).Field(0).String()
			if i < ref_v.Len()-1 {
				ids_str += ","
			}
		}
	case User:
		ids_str = ref_v.Field(0).String()
	case string:
		ids_str = ref_v.String()
	default:
		panic("Wrong argument. fetchIds() can take in  String, []String, []User or User.")
	}
	return ids_str
}
