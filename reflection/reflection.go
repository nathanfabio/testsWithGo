package reflection

import "reflect"

func Browse(x interface{}, fn func(input string)) {
	value := getValue(x)

	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			Browse(value.Index(i).Interface(),fn)
		}
		return
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Browse(field.Interface(), fn)
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