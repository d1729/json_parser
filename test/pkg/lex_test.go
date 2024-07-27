package pkg

import (
	"json_parser/pkg"
	"reflect"
	"testing"
)

func TestLexString(t *testing.T) {
	s1, s2 := pkg.LexString("\"Hello\": \"Some\"")
	if *s1 != "Hello" {
		t.Errorf("Wanted Hello, got %s", *s1)
	}
	if s2 != ": \"Some\"" {
		t.Errorf("Wanted : \"Some\", got %s", s2)
	}
}

func TestLexStringWithNilParsed(t *testing.T) {
	s1, s2 := pkg.LexString("123: \"Some\"")
	if s1 != nil {
		t.Errorf("Wanted Hello, got %s", *s1)
	}
	if s2 != "123: \"Some\"" {
		t.Errorf("Wanted 123: \"Some\", got %s", s2)
	}
}

func TestLexNumber(t *testing.T) {
	num1, s2 := pkg.LexNumber("123: \"Hi\"")
	if *num1 != 123 {
		t.Errorf("Wanted 123, got %d", num1)
	}
	if s2 != ": \"Hi\"" {
		t.Errorf("Wanted : \"Hi\", got %s", s2)
	}
}

func TestLexNumberWithNilParsed(t *testing.T) {
	num1, s2 := pkg.LexNumber("\"Nope\": \"Hi\"")
	if num1 != nil {
		t.Errorf("Wanted nil, got %d", num1)
	}
	if s2 != "\"Nope\": \"Hi\"" {
		t.Errorf("Wanted \"Nope\": \"Hi\", got %s", s2)
	}
}

func TestLexNumberWithFloat(t *testing.T) {
	num1, s2 := pkg.LexNumber("123.12: \"Hi\"")
	if *num1 != 123.12 {
		t.Errorf("Wanted 123.12, got %f", *num1)
	}
	if s2 != ": \"Hi\"" {
		t.Errorf("Wanted: \"Hi\", got %s", s2)
	}
}

func TestLexNumberWithWrongFloat(t *testing.T) {
	num1, s2 := pkg.LexNumber("123.: \"Hi\"")
	if *num1 != 123 {
		t.Errorf("Wanted 123., got %f", *num1)
	}
	if s2 != ": \"Hi\"" {
		t.Errorf("Wanted: \"Hi\", got %s", s2)
	}
}

func TestLexBoolTrue(t *testing.T) {
	val, s := pkg.LexBool("true, 123")
	if *val != true {
		t.Errorf("Wanted true, got %v", *val)
	}
	if s != ", 123" {
		t.Errorf("Wanted , 123, got %v", s)
	}
}

func TestLexBoolFalse(t *testing.T) {
	val, s := pkg.LexBool("false, 123")
	if *val != false {
		t.Errorf("Wanted false, got %v", *val)
	}
	if s != ", 123" {
		t.Errorf("Wanted , 123, got %v", s)
	}
}

func TestLexBoolWrong(t *testing.T) {
	val, s := pkg.LexBool("123, 245")
	if val != nil {
		t.Errorf("Wanted nil, got %v", *val)
	}
	if s != "123, 245" {
		t.Errorf("Wanted 123, 245, got %v", s)
	}
}

func TestLexNull(t *testing.T) {
	val, s := pkg.LexNull("null, 123")
	if *val != true {
		t.Errorf("Wanted true, got %v", *val)
	}
	if s != ", 123" {
		t.Errorf("Wanted , 123, got %v", s)
	}
}

func TestLexNullWithNilParsed(t *testing.T) {
	val, s := pkg.LexNull(", 123")
	if val != nil {
		t.Errorf("Wanted nil, got %v", *val)
	}
	if s != ", 123" {
		t.Errorf("Wanted , 123, got %v", s)
	}
}

func TestLex(t *testing.T) {
	tests := []struct {
		json string
		vals []interface{}
	}{
		{"{\"deb\":1.2,}", []interface{}{"{", "deb", ":", 1.2, ",", "}"}},
		{"{\"foo\":[1,2,{\"bar\":2}]}", []interface{}{"{", "foo", ":", "[", 1, ",", 2, ",", "{", "bar", ":", 2, "}", "]", "}"}},
		{"{}", []interface{}{"{", "}"}},
	}

	for _, test := range tests {
		res := pkg.Lex(test.json)
		if len(res) != len(test.vals) {
			t.Errorf("Wanted %d, got %d", len(test.vals), len(res))
		}
		for i := range res {
			if res[i] != test.vals[i] {
				t.Errorf("%v, %v", reflect.TypeOf(test.vals[i]), reflect.TypeOf(res[i]))
				t.Errorf("Following values don't match %v, %v", test.vals[i], res[i])
			}
		}
	}

}
