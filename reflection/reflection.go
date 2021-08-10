package reflectioness

import (
	"reflect"
)

func walk(x interface{}, function func(input string)) {
	value := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), function)
	}

	switch value.Kind() {
	case reflect.String:
		function(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := value.Recv(); ok; v, ok = value.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := value.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
