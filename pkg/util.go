package pkg

import "reflect"

func IsZeroValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	case reflect.String:
		return val.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return val.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Bool:
		return !val.Bool()
	}

	return false
}
