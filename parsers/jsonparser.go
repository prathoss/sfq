package parsers

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

var depthIncreasers = []rune{'{', '['}
var depthReducers = []rune{'}', ']'}
var valueClosers = append(depthReducers, ',')

func rContains(ru rune, rs []rune) bool {
	for _, r := range rs {
		if r == ru {
			return true
		}
	}
	return false
}

type jsonParser struct {
	stringBuilder   strings.Builder
	isKey           bool
	isValue         bool
	insideString    bool
	read            bool
	depth           int
	currentKeyDepth int
}

func (p *jsonParser) Parse(reader io.Reader, keyHandler func(string, int) bool, valueHandler func(string), otherSymbolHandler func(rune)) error {
	if keyHandler == nil {
		return fmt.Errorf("keyHandler function can not be nil")
	}

	if valueHandler == nil {
		return fmt.Errorf("valueHandler function can not be nil")
	}

	decoder := json.NewDecoder(reader)
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		fmt.Println(token)
	}
	return nil
}
