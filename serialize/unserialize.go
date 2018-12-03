package serialize

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/techoner/gophp/utils"
)

const UNSERIALIZABLE_OBJECT_MAX_LEN = 10 * 1024 * 1024 * 1024

func UnMarshal(data []byte) (interface{}, error) {
	reader := bytes.NewReader(data)
	return unMarshalByReader(reader)
}

func unMarshalByReader(reader *bytes.Reader) (interface{}, error) {

	for {

		if token, _, err := reader.ReadRune(); err == nil {
			switch token {
			default:
				return nil, fmt.Errorf("UnMarshal: Unknown token %#U", token)
			case 'N':
				return unMarshalNil(reader)
			case 'b':
				return unMarshalBool(reader)
			case 'i':
				return unMarshalNumber(reader, false)
			case 'd':
				return unMarshalNumber(reader, true)
			case 's':
				return unMarshalString(reader, true)
			case 'a':
				return unMarshalArray(reader)
				// case 'O':

				// case 'C':

				// case 'R', 'r':

				// case 'x':

			}
		}
		return nil, nil
	}

}

func unMarshalNil(reader *bytes.Reader) (interface{}, error) {
	expect(reader, ';')

	return nil, nil
}

func unMarshalBool(reader *bytes.Reader) (interface{}, error) {
	var (
		raw rune
		err error
	)
	err = expect(reader, ':')
	if err != nil {
		return nil, err
	}

	if raw, _, err = reader.ReadRune(); err != nil {
		return nil, fmt.Errorf("UnMarshal: Error while reading bool value: %v", err)
	}

	err = expect(reader, ';')
	if err != nil {
		return nil, err
	}
	return raw == '1', nil
}

func unMarshalNumber(reader *bytes.Reader, isFloat bool) (interface{}, error) {
	var (
		raw string
		err error
		val interface{}
	)
	err = expect(reader, ':')
	if err != nil {
		return nil, err
	}

	if raw, err = readUntil(reader, ';'); err != nil {
		return nil, fmt.Errorf("UnMarshal: Error while reading number value: %v", err)
	} else {
		if isFloat {
			if val, err = strconv.ParseFloat(raw, 64); err != nil {
				return nil, fmt.Errorf("UnMarshal: Unable to convert %s to float: %v", raw, err)
			}
		} else {
			if val, err = strconv.Atoi(raw); err != nil {
				return nil, fmt.Errorf("UnMarshal: Unable to convert %s to int: %v", raw, err)
			}
		}
	}

	return val, nil
}

func unMarshalString(reader *bytes.Reader, isFinal bool) (interface{}, error) {
	var (
		err     error
		val     interface{}
		strLen  int
		readLen int
	)

	strLen, err = readLength(reader)

	err = expect(reader, '"')
	if err != nil {
		return nil, err
	}

	if strLen > 0 {
		buf := make([]byte, strLen, strLen)
		if readLen, err = reader.Read(buf); err != nil {
			return nil, fmt.Errorf("UnMarshal: Error while reading string value: %v", err)
		} else {
			if readLen != strLen {
				return nil, fmt.Errorf("UnMarshal: Unable to read string. Expected %d but have got %d bytes", strLen, readLen)
			} else {
				val = string(buf)
			}
		}
	}

	err = expect(reader, '"')
	if err != nil {
		return nil, err
	}
	if isFinal {
		err = expect(reader, ';')
		if err != nil {
			return nil, err
		}
	}
	return val, nil
}

func unMarshalArray(reader *bytes.Reader) (interface{}, error) {
	var arrLen int
	var err error
	val := make(map[string]interface{})

	arrLen, err = readLength(reader)

	if err != nil {
		return nil, err
	}
	err = expect(reader, '{')
	if err != nil {
		return nil, err
	}
	indexLen := 0
	for i := 0; i < arrLen; i++ {
		k, err := unMarshalByReader(reader)
		if err != nil {
			return nil, err
		}
		v, err := unMarshalByReader(reader)
		if err != nil {
			return nil, err
		}

		// if errKey == nil && errVal == nil {
		// val[k] = v
		switch t := k.(type) {
		default:
			return nil, fmt.Errorf("UnMarshal: Unexpected key type %T", t)
		case string:
			stringKey, _ := k.(string)
			val[stringKey] = v
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			// intKey, _ := k.(int)
			// val[strconv.Itoa(intKey)] = v
			stringKey, _ := utils.NumericalToString(k)
			val[stringKey] = v

			// stringI, _ := utils.NumericalToString(i)
			if i == k {
				indexLen++
			}

		}
		// } else {
		// 	return nil, fmt.Errorf("UnMarshal: Error while reading key or(and) value of array")
		// }
	}

	err = expect(reader, '}')
	if err != nil {
		return nil, err
	}

	if indexLen == arrLen {
		var slice []interface{}
		for _, row := range val {
			slice = append(slice, row)
		}
		return slice, nil
	}

	return val, nil
}

func expect(reader *bytes.Reader, expected rune) error {
	if token, _, err := reader.ReadRune(); err != nil {
		return fmt.Errorf("UnMarshal: Error while reading expected rune %#U: %v", expected, err)
	} else if token != expected {
		return fmt.Errorf("UnMarshal: Expected %#U but have got %#U", expected, token)
	}
	return nil
}

func readUntil(reader *bytes.Reader, stop rune) (string, error) {
	var (
		token rune
		err   error
	)
	buf := bytes.NewBuffer([]byte{})

	for {
		if token, _, err = reader.ReadRune(); err != nil || token == stop {
			break
		} else {
			buf.WriteRune(token)
		}
	}

	return buf.String(), err
}

func readLength(reader *bytes.Reader) (int, error) {
	var (
		raw string
		err error
		val int
	)
	err = expect(reader, ':')
	if err != nil {
		return 0, err
	}

	if raw, err = readUntil(reader, ':'); err != nil {
		return 0, fmt.Errorf("UnMarshal: Error while reading lenght of value: %v", err)
	} else {
		if val, err = strconv.Atoi(raw); err != nil {
			return 0, fmt.Errorf("UnMarshal: Unable to convert %s to int: %v", raw, err)
		} else if val > UNSERIALIZABLE_OBJECT_MAX_LEN {
			return 0, fmt.Errorf("UnMarshal: Unserializable object length looks too big(%d). If you are sure you wanna unserialise it, please increase UNSERIALIZABLE_OBJECT_MAX_LEN const", val)
			val = 0
		}
	}
	return val, nil
}
