package pkg

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func LexString(str string) (*string, string) {
	jsonString := ""
	start := 0

	if string(str[0]) == JsonQuote {
		start = 1
	} else {
		return nil, str
	}

	for i := start; i < len(str); i++ {
		c := str[i]
		if string(c) == JsonQuote {
			return &jsonString, str[i+1:]
		} else {
			jsonString += string(c)
		}
	}

	panic("Expected end of line quote")
}

func LexNumber(str string) (*float64, string) {
	jsonNumber := ""

	var numbers []string
	for i := range 10 {
		numbers = append(numbers, strconv.Itoa(i))
	}
	numbers = append(numbers, "-", "e", ".")

	for _, c := range str {
		if slices.Contains(numbers, string(c)) {
			jsonNumber += string(c)
		} else {
			break
		}
	}

	if len(jsonNumber) == 0 {
		return nil, str
	}

	if strings.ContainsRune(jsonNumber, '.') {
		//if strings.IndexRune(jsonNumber, '.') == len(jsonNumber)-1 {
		//	panic("Number not formed properly")
		//} else {
		num, err := strconv.ParseFloat(jsonNumber, 64)
		if err != nil {
			panic("Number not formed properly")
		}
		return &num, str[len(jsonNumber):]
		//}
	}

	num, err := strconv.Atoi(jsonNumber)
	if err != nil {
		panic("Number not formed properly")
	}
	floatVal := float64(num)
	return &floatVal, str[len(jsonNumber):]
}

func LexBool(str string) (*bool, string) {
	stringLen := len(str)

	if stringLen >= TrueLen && str[:TrueLen] == "true" {
		val := true
		return &val, str[TrueLen:]
	}
	if stringLen >= FalseLen && str[:FalseLen] == "false" {
		val := false
		return &val, str[FalseLen:]
	}

	return nil, str
}

func LexNull(str string) (*bool, string) {
	stringLen := len(str)

	if stringLen >= NullLen && str[:NullLen] == "null" {
		val := true
		return &val, str[NullLen:]
	}
	return nil, str
}

func Lex(str string) []interface{} {
	var tokens []interface{}

	for i := 0; i < len(str); {
		jsonString, newStr := LexString(str)
		if jsonString != nil {
			tokens = append(tokens, *jsonString)
			str = newStr
			continue
		}

		jsonNumber, newStr := LexNumber(str)
		if jsonNumber != nil {
			if *jsonNumber == math.Floor(*jsonNumber) {
				tokens = append(tokens, int(*jsonNumber))
			} else {
				tokens = append(tokens, *jsonNumber)
			}
			str = newStr
			continue
		}

		jsonBool, newStr := LexBool(str)
		if jsonBool != nil {
			tokens = append(tokens, *jsonBool)
			str = newStr
			continue
		}

		jsonNull, newStr := LexNull(str)
		if jsonNull != nil {
			tokens = append(tokens, *jsonNull)
			str = newStr
			continue
		}

		_, existsInWhiteSpaces := JsonWhiteSpaces[rune(str[0])]
		_, existsInJsonSyntax := JsonSyntax[rune(str[0])]
		if existsInWhiteSpaces {
			str = str[1:]
		} else if existsInJsonSyntax {
			tokens = append(tokens, string(str[0]))
			str = str[1:]
		} else {
			panic(fmt.Sprintf("Unexpected token %c", rune(str[0])))
		}

	}

	return tokens

}
