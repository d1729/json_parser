package pkg

import (
	"fmt"
	"reflect"
)

type ParseResult struct {
	JsonArray  []interface{}
	JsonObject map[interface{}]interface{}
	Token      interface{}
}

func parseArray(tokens []interface{}) (ParseResult, []interface{}) {
	res := ParseResult{}
	var jsonArray []interface{}

	t := tokens[0]
	if t == JsonRightBracket {
		res.JsonArray = jsonArray
		return res, tokens[1:]
	}

	for {
		newToken, tempTokens := Parse(tokens)
		json := newToken.Token
		if newToken.JsonArray != nil {
			json = newToken.JsonArray
		} else if newToken.JsonObject != nil {
			json = newToken.JsonObject
		}
		jsonArray = append(jsonArray, json)

		t = tempTokens[0]
		if t == JsonRightBracket {
			res.JsonArray = jsonArray
			return res, tempTokens[1:]
		} else if t != JsonComma {
			panic("Expected a comma ")
		} else {
			tokens = tempTokens[1:]
		}
	}

}

func parseObject(tokens []interface{}) (ParseResult, []interface{}) {
	res := ParseResult{}
	var jsonObject = make(map[interface{}]interface{})

	t := tokens[0]
	if t == JsonRightBrace {
		res.JsonObject = jsonObject
		return res, tokens[1:]
	}

	for {
		jsonKey := tokens[0]
		if reflect.TypeOf(jsonKey) == reflect.TypeOf("abc") || reflect.TypeOf(jsonKey) == reflect.TypeOf(1) {
			tokens = tokens[1:]
		} else {
			panic(fmt.Sprintf("Expected string or integer key, got: %v", jsonKey))
		}

		if tokens[0] != JsonColon {
			panic(fmt.Sprintf("Expected a colon after key in object, got %v", tokens[0]))
		}

		newToken, tempTokens := Parse(tokens[1:])
		jsonValue := newToken.Token
		if newToken.JsonArray != nil {
			jsonValue = newToken.JsonArray
		} else if newToken.JsonObject != nil {
			jsonValue = newToken.JsonObject
		}
		jsonObject[jsonKey] = jsonValue

		t = tempTokens[0]
		if t == JsonRightBrace {
			res.JsonObject = jsonObject
			return res, tempTokens[1:]
		} else if t != JsonComma {
			panic(fmt.Sprintf("Expected a comma after pair in object got %v", t))
		}

		tokens = tempTokens[1:]
	}
}

func Parse(tokens []interface{}) (ParseResult, []interface{}) {
	//l := log.Default()
	//l.Print(tokens)
	t := tokens[0]

	if t == JsonLeftBracket {
		return parseArray(tokens[1:])
	} else if t == JsonLeftBrace {
		return parseObject(tokens[1:])
	} else {
		return ParseResult{Token: t}, tokens[1:]
	}
}
