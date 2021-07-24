package reflectioness

import (
	"reflect"
)

func walk(x interface{}, function func(input string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			function(field.String())
		case reflect.Struct:
			walk(field.Interface(), function)
		}
	}
}
