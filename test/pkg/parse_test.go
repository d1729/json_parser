package pkg

import (
	"json_parser/pkg"
	"testing"
)

func TestParse(t *testing.T) {
	parseResult, rest := pkg.Parse([]interface{}{"{", "foo", ":", "[", 1, ",", 2, ",", "{", "bar", ":", 2, "}", "]", "}"})
	//t.Log(parseResult, rest)
	if len(parseResult.JsonObject) != 1 {
		t.Errorf("Length of parseResult.JsonObject should be 1, but it is %d", len(parseResult.JsonObject))
	}
	if len(rest) != 0 {
		t.Errorf("Length of rest should be 0, but it is %d", len(rest))
	}
}
