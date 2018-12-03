package gophp

import (
	"fmt"
	"reflect"
	"strings"
)

// Implode Join array elements with a string
func Implode(separator string, array []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
}

// IsArray Finds whether a variable is an array
func IsArray(v interface{}) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	// case reflect.Map:
	// 	return true
	case reflect.Slice:
		return true
	default:

	}

	return false
}

// ArrayKeys Return all the keys or a subset of the keys of an array
func ArrayKeys(array map[string]interface{}) []string {
	var keys []string
	for k := range array {
		keys = append(keys, k)
	}
	return keys
}
