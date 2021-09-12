package v1

import (
	"reflect"
)

func Walk(a interface{}, fn func(string)) {
	val := reflect.ValueOf(a)
	
	field := val.Field(0)

	fn(field.String())
}