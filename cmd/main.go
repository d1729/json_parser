package cmd

import (
	"json_parser/pkg"
)

func JsonToObj(jsonStr string) (pkg.ParseResult, []interface{}, int) {
	tokens := pkg.Lex(jsonStr)
	res, rest := pkg.Parse(tokens)
	if len(rest) != 0 {
		return res, rest, -1
	}
	return res, rest, 0
}
