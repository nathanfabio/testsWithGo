package reflection

import "reflect"

func Browse(x interface{}, fn func(input string)) {
	value := getValue(x)

	brwseValue := func(value reflect.Value) {
		Browse(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			brwseValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			brwseValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			brwseValue(value.MapIndex(key))
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