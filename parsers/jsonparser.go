package parsers

import (
	"fmt"
	"strings"
	"unicode"
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

func (p *jsonParser) Parse(fileName string, keyHandler func(string, int) bool, valueHandler func(string), otherSymbolHandler func(rune)) error {
	if keyHandler == nil {
		return fmt.Errorf("keyHandler function can not be nil")
	}

	if valueHandler == nil {
		return fmt.Errorf("valueHandler function can not be nil")
	}

	isOtherSymbolApplicable := func() bool {
		return otherSymbolHandler != nil && !p.read
	}

	return readFile(fileName, func(text string) {
		for _, r := range text {
			if rContains(r, depthIncreasers) {
				p.depth++
				p.isValue = false
				continue
			}
			if (unicode.IsSpace(r) || r == ':') && !p.insideString {
				continue
			}
			//if isKey and isValue are false => looking for begining of a key
			if !p.isKey && !p.isValue && string(r) == "\"" {
				p.isKey = !p.isKey
				p.read = true
				continue
			}
			//reading key, when " found must mean it is end of the key
			if p.isKey && string(r) == "\"" {
				p.read = keyHandler(p.stringBuilder.String(), p.depth)
				if p.read {
					p.currentKeyDepth = p.depth
				}
				p.stringBuilder.Reset()
				p.isKey = !p.isKey
				p.isValue = !p.isValue
				continue
			}
			//reading value, when value closer found must mean end of value
			if p.isValue && rContains(r, valueClosers) && p.read && p.depth == p.currentKeyDepth {
				valueHandler(p.stringBuilder.String())
				p.stringBuilder.Reset()
				p.isValue = !p.isValue
				p.read = true
				continue
			}
			if rContains(r, depthReducers) {
				p.depth--
			}
			if p.read && p.depth >= 0 {
				p.stringBuilder.WriteRune(r)
			}
			if isOtherSymbolApplicable() {
				otherSymbolHandler(r)
			}
			if p.depth == 0 && rContains(r, valueClosers) {
				p.isValue = false
			}
		}
	})
}
