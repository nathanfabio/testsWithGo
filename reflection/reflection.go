package reflection

import "reflect"

func Browse(x interface{}, fn func(input string)) {
	value := reflect.ValueOf(x)
	field := value.Field(0)
	fn(field.String())
}