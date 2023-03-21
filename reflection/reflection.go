package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// you're passing in a value whose type is `interface{}` aka "any"
	// get a concrete value stored in the interface `x`
	val := reflect.ValueOf(x)

	// grab the 0th field of the struct
	// and: it only works for structs
	field := val.Field(0)

	// we're making the assumption that the struct field is a string
	fn(field.String())
}
