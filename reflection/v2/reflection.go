package v2

import (
	"reflect"
)

func Walk(a interface{}, fn func(string)) {
	val := GetValue(a)
	numberOfValues := 0
	var getField func(int) reflect.Value
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
        getField = val.Index
		// for i:=0; i< val.Len(); i++ {
        //    Walk(val.Index(i).Interface(), fn)
		// }
	case reflect.String:
        fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
		// for i:=0; i<val.NumField(); i++ {
        //     Walk(val.Field(i).Interface(), fn)
        // }
	case reflect.Map:
        for _, key := range val.MapKeys() {
            Walk(val.MapIndex(key).Interface(), fn)
        }
	}
	for i:=0; i< numberOfValues; i++ {
        Walk(getField(i).Interface(), fn)
    }

}

func GetValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
	return val
}