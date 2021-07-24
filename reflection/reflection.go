package reflectioness

import (
	"reflect"
)

func walk(x interface{}, function func(input string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		function(field.String())
	}
}
