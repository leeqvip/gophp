package serialize

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/techleeone/gophp"
)

func Marshal(value interface{}) ([]byte, error) {

	if value == nil {
		return MarshalNil(), nil
	}

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Bool:
		return MarshalBool(value.(bool)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return MarshalNumber(value), nil
	case reflect.String:
		return MarshalString(value.(string)), nil
	case reflect.Map:
		return MarshalMap(value)
	case reflect.Slice:
		return MarshalSlice(value)
	case reflect.Struct:
	default:
		return nil, fmt.Errorf("Marshal: Unknown type %T with value %#v", t, value)
	}

	return nil, nil
}

func MarshalNil() []byte {
	return []byte("N;")
}

func MarshalBool(value bool) []byte {
	if value {
		return []byte("b:1;")
	}

	return []byte("b:0;")
}

func MarshalNumber(value interface{}) []byte {
	var val string

	isFloat := false

	switch value.(type) {
	default:
		val = "0"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		val, _ = gophp.NumericalToString(value)
	case float32, float64:
		val, _ = gophp.NumericalToString(value)
		isFloat = true
	}

	if isFloat {
		return []byte("d:" + val + ";")

	} else {
		return []byte("i:" + val + ";")
	}
}

func MarshalString(value string) []byte {
	return []byte(fmt.Sprintf("s:%d:\"%s\";", len(value), value))
}

func MarshalMap(value interface{}) ([]byte, error) {

	s := reflect.ValueOf(value)

	// Go randomises maps. To be able to test this we need to make sure the
	// map keys always come out in the same order. So we sort them first.
	mapKeys := s.MapKeys()
	sort.Slice(mapKeys, func(i, j int) bool {
		return gophp.LessValue(mapKeys[i], mapKeys[j])
	})

	var buffer bytes.Buffer
	for _, mapKey := range mapKeys {
		m, err := Marshal(mapKey.Interface())
		if err != nil {
			return nil, err
		}

		buffer.Write(m)

		m, err = Marshal(s.MapIndex(mapKey).Interface())
		if err != nil {
			return nil, err
		}

		buffer.Write(m)
	}

	return []byte(fmt.Sprintf("a:%d:{%s}", s.Len(), buffer.String())), nil
}

func MarshalSlice(value interface{}) ([]byte, error) {
	s := reflect.ValueOf(value)

	var buffer bytes.Buffer
	for i := 0; i < s.Len(); i++ {
		m, err := Marshal(i)
		if err != nil {
			return nil, err
		}

		buffer.Write(m)

		m, err = Marshal(s.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		buffer.Write(m)
	}

	return []byte(fmt.Sprintf("a:%d:{%s}", s.Len(), buffer.String())), nil
}

func MarshalStruct(input interface{}) ([]byte, error) {
	value := reflect.ValueOf(input)
	typeOfValue := value.Type()

	// Some of the fields in the struct may not be visible (unexported). We
	// need to make sure we count all the visible ones for the final result.
	visibleFieldCount := 0

	var buffer bytes.Buffer
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)

		if !f.CanInterface() {
			// This is an unexported field, we cannot read it.
			continue
		}

		visibleFieldCount++

		// Note: since we can only export fields that are public (start
		// with an uppercase letter) we must change it to lower case. If
		// you really do want it to be upper case you will have to wait
		// for when tags are supported on individual fields.
		fieldName := gophp.LowerCaseFirstLetter(typeOfValue.Field(i).Name)
		buffer.Write(MarshalString(fieldName))

		m, err := Marshal(f.Interface())
		if err != nil {
			return nil, err
		}

		buffer.Write(m)
	}

	className := reflect.ValueOf(input).Type().Name()

	return []byte(fmt.Sprintf("O:%d:\"%s\":%d:{%s}", len(className),
		className, visibleFieldCount, buffer.String())), nil
}
