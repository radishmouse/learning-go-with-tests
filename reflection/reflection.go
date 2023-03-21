package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// you're passing in a value whose type is `interface{}` aka "any"
	// get a concrete value stored in the interface `x`
	val := reflect.ValueOf(x)

	// now...I'm going to receive more than one field.
	// how do I count the number of fields?
	for i := 0; i < val.NumField(); i++ {
		// get the ith field from `val`, which must be a struct (else err)
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}

}
