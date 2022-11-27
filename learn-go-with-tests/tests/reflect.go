package tests

import "reflect"

func getValueOf(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, fn func(input string)) {
	val := getValueOf(x)
	walkValue := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			// walk(val.Field(i).Interface(), fn)
			walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			// walk(val.Index(i).Interface(), fn)
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, k := range val.MapKeys() {
			// walk(val.MapIndex(k).Interface(), fn)
			walkValue(val.MapIndex(k))
		}
	}
}
