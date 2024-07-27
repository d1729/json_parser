package cmd

import (
	"json_parser/cmd"
	"json_parser/pkg"
	"testing"
)

func TestJsonToObj(t *testing.T) {
	tests := []struct {
		input  string
		res    pkg.ParseResult
		rest   []interface{}
		signal int
	}{
		{"{}", pkg.ParseResult{JsonArray: []interface{}{}, JsonObject: map[interface{}]interface{}{}, Token: []interface{}{}}, []interface{}{}, 0},
		{"{\n  \"key\": \"value\",\n  \"key2\": \"value\"\n}", pkg.ParseResult{JsonArray: []interface{}{}}, []interface{}{}, 0},
	}
	for _, test := range tests {
		_, _, signal := cmd.JsonToObj(test.input)
		if signal != test.signal {
			t.Errorf("Signal = %d, want %d", signal, test.signal)
		}
	}

}
