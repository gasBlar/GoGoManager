package utils

import "reflect"

func IsValidField(structType reflect.Type, fieldName string) bool {
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("json")
		if fieldName == tag {
			return true
		}
	}
	return false
}
