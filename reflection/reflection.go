package reflectioness

import (
	"reflect"
)

func walk(x interface{}, function func(input string)) {
	value := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		function(value.String())
	case reflect.Struct:
		numberOfValues = value.NumField()
		getField = value.Field
	case reflect.Slice:
		numberOfValues = value.Len()
		getField = value.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), function)
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
