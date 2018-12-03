package serialize

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/techoner/gophp/utils"
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
		val, _ = utils.NumericalToString(value)
	case float32, float64:
		val, _ = utils.NumericalToString(value)
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

	mapKeys := s.MapKeys()
	sort.Slice(mapKeys, func(i, j int) bool {
		return utils.LessValue(mapKeys[i], mapKeys[j])
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
