package serialize

import "testing"

func TestMarshalNil(t *testing.T) {

	out := MarshalNil()

	if string(out) != "N;" {
		t.Errorf("Nil value marshaled incorrectly, have got %q\n", out)
	}

}

func TestMarshalBoolTrue(t *testing.T) {

	out := MarshalBool(true)

	if string(out) != "b:1;" {
		t.Errorf("Bool value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalBoolFalse(t *testing.T) {

	out := MarshalBool(false)

	if string(out) != "b:0;" {
		t.Errorf("Bool value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberInt(t *testing.T) {

	out := MarshalNumber(10)

	if string(out) != "i:10;" {
		t.Errorf("Int value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberInt8(t *testing.T) {

	out := MarshalNumber(int8(10))

	if string(out) != "i:10;" {
		t.Errorf("Int8 value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberInt16(t *testing.T) {

	out := MarshalNumber(int16(12))

	if string(out) != "i:12;" {
		t.Errorf("Int16 value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberInt32(t *testing.T) {

	out := MarshalNumber(int32(123456))

	if string(out) != "i:123456;" {
		t.Errorf("Int32 value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberInt64(t *testing.T) {

	out := MarshalNumber(int64(12345678910))

	if string(out) != "i:12345678910;" {
		t.Errorf("Int64 value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberUInt(t *testing.T) {

	out := MarshalNumber(uint(12345678910))

	if string(out) != "i:12345678910;" {
		t.Errorf("Uint value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberFloat32(t *testing.T) {

	var input float32
	input = 1.25
	out := MarshalNumber(input)

	if string(out) != "d:1.25;" {
		t.Errorf("float32 value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalNumberFloat64(t *testing.T) {

	var input float64
	input = 1.25
	out := MarshalNumber(input)

	if string(out) != "d:1.25;" {
		t.Errorf("float64 value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalString(t *testing.T) {

	var input string
	input = "PHP is a popular general-purpose scripting language that is especially suited to web development. PHP是一种流行的通用脚本语言，特别适用于Web开发。"
	out := MarshalString(input)

	if string(out) != `s:167:"PHP is a popular general-purpose scripting language that is especially suited to web development. PHP是一种流行的通用脚本语言，特别适用于Web开发。";` {
		t.Errorf("String value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalMap(t *testing.T) {

	input := map[interface{}]interface{}{
		"language":    "PHP",
		"description": "a popular general-purpose scripting language",
	}

	out, err := MarshalMap(input)
	if err != nil {
		panic(err)
	}

	if string(out) != `a:2:{s:8:"language";s:3:"PHP";s:11:"description";s:44:"a popular general-purpose scripting language";}` {
		t.Errorf("Map value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshalSlice(t *testing.T) {
	var input []interface{}

	input = append(input, map[interface{}]interface{}{
		"language":    "PHP",
		"description": "a popular general-purpose scripting language",
	})

	out, err := MarshalSlice(input)
	if err != nil {
		panic(err)
	}

	if string(out) != `a:1:{i:0;a:2:{s:8:"language";s:3:"PHP";s:11:"description";s:44:"a popular general-purpose scripting language";}}` {
		t.Errorf("Slice value marshaled incorrectly, have got %q\n", out)
	}
}

func TestMarshal(t *testing.T) {
	var input []interface{}

	input = append(input, map[interface{}]interface{}{
		"display_url":           "/group/6616191620721148423/",
		"title":                 "一段不幸的婚姻害死丈夫 妻子法庭上向家人下跪赎罪 婆婆情绪失控",
		"pc_image_url":          "https://p99.pstatp.com/list/300x170/pgc-image/15404520275419b14c54a1a",
		"comment_count":         520,
		"video_play_count":      951011,
		"video_duration_format": "15:00",
		"video_duration":        900,
	})
	input = append(input, map[interface{}]interface{}{
		"display_url":           "/group/6607288769219396103/",
		"title":                 "重庆美女司机学车，教练说，你入党了么？笑翻了",
		"pc_image_url":          "https://p3.pstatp.com/list/300x170/cc390008433eac69241b",
		"comment_count":         76,
		"video_play_count":      1587177,
		"video_duration_format": "03:55",
		"video_duration":        235,
	})

	_, err := Marshal(input)
	if err != nil {
		panic(err)
	}
}
