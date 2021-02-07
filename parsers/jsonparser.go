package parsers

import (
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
	stringBuilder strings.Builder
	isKey         bool
	isValue       bool
	read          bool
	depth         int
}

func (p *jsonParser) Parse(fileName string, keyHandler func(string) bool, valueHandler func(string)) error {
	return readFile(fileName, func(text string) {
		for _, r := range text {
			if rContains(r, depthIncreasers) {
				p.depth++
			}
			//if isKey and isValue are false => looking for begining of a key
			if !p.isKey && !p.isValue && string(r) == "\"" {
				p.isKey = !p.isKey
				p.read = true
				continue
			}
			//reading key, when " found must mean it is end of the key
			if p.isKey && string(r) == "\"" {
				p.read = keyHandler(p.stringBuilder.String())
				p.stringBuilder.Reset()
				p.isKey = !p.isKey
				p.isValue = !p.isValue
				continue
			}
			//reading value, when value closer found must mean end of value
			if p.isValue && rContains(r, valueClosers) && p.read && p.depth == 0 {
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
			if p.depth == 0 && rContains(r, valueClosers) {
				p.isValue = false
			}
		}
	})
}
