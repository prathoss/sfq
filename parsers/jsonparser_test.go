package parsers

import (
	"io/ioutil"
	"testing"
)

func TestParseJSON(t *testing.T) {
	json, err := ioutil.TempFile("", "jsonParse")
	if err != nil {
		t.Fatal(err)
	}
	json.WriteString(`
	{
		"key1": "value1",
		"key2": "value2"
	}`)

	keys := []string{}
	values := []string{}

	parser := NewJsonParser()
	parser.Parse(json, func(s string, i int) KeyAction {
		keys = append(keys, s)
		return ReturnAction
	}, func(s string) {
		values = append(values, s)
	}, nil)

	if keys[0] != "key1" {
		t.Error("first key should be 'key1'")
	}
	if keys[1] != "key2" {
		t.Error("second key should be 'key2'")
	}
	if values[0] != "value1" {
		t.Error("first value should be 'value1'")
	}
	if values[1] != "value2" {
		t.Error("first value should be 'value2'")
	}
}
