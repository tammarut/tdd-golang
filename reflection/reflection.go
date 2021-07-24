package reflectioness

import (
	"reflect"
)

func walk(x interface{}, function func(input string)) {
	value := getValue(x)

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), function)
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), function)
		}
	case reflect.String:
		function(value.String())
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
