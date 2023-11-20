package utils

import (
	"reflect"
)

func Has(object interface{}, key string) bool {
	value := reflect.ValueOf(object)
	field := value.FieldByName(key)

	return !field.IsZero()
}
