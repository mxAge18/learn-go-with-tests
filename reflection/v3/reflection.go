package v3

import (
	"reflect"
)

func Walk(a interface{}, fn func(string)) {
	val := GetValue(a)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i:=0; i< val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.String:
        fn(val.String())
	case reflect.Struct:
		for i:=0; i<val.NumField(); i++ {
			walkValue(val.Field(i))
        }
	case reflect.Map: // Go 中的 map 不能保证顺序一致。因此，你的测试有时会失败，因为我们断言对 fn 的调用是以特定的顺序完成的。
        for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
        }
	}
}

func GetValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
	return val
}