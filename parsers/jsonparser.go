package parsers

import (
	"strings"
)

type jsonParser struct {
	stringBuilder strings.Builder
	isKey         bool
	isValue       bool
	read          bool
}

func (p *jsonParser) Parse(fileName string, keyHandler func(string) bool, valueHandler func(string)) error {
	return readFile(fileName, func(text string) {
		for _, r := range text {
			if !p.isKey && !p.isValue && string(r) == "\"" {
				p.isKey = !p.isKey
				p.read = true
				continue
			}
			if p.isKey && string(r) == "\"" {
				p.read = keyHandler(p.stringBuilder.String())
				p.stringBuilder.Reset()
				p.isKey = !p.isKey
				p.isValue = !p.isValue
				continue
			}
			if p.isValue && (string(r) == "," || string(r) == "\n") {
				valueHandler(p.stringBuilder.String())
				p.stringBuilder.Reset()
				p.isValue = !p.isValue
				p.read = true
				continue
			}
			if p.read {
				p.stringBuilder.WriteRune(r)
			}
		}
	})
}
