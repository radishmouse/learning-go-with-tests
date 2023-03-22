package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		// :) I get to use the combo assign/eval trick:
		switch field := val.Field(i); field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	// you're passing in a value whose type is `interface{}` aka "any"
	// get a concrete value stored in the interface `x`
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		// Here's how you dereference to get the concrete value
		val = val.Elem()
	}

	return val
}
