package parsers

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func NewJsonParser() jsonParser {
	return jsonParser{}
}

type jsonParser struct {
	stringBuilder   strings.Builder
	isKey           bool
	controlKeyDepth int
	immersion       []immersionType
	keyAction       KeyAction
}

func (p *jsonParser) Parse(reader io.ReadSeekCloser, keyHandler func(string, int) KeyAction, valueHandler func(string), otherSymbolHandler func(rune)) error {
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
		decoder.InputOffset()
		switch token {
		case "{":
			p.immersion = append(p.immersion, imObject)
			p.isKey = true
		case "[":
			p.immersion = append(p.immersion, imArray)
			p.isKey = false
		case "}", "]":
			p.immersion = p.immersion[:len(p.immersion)-1]
			p.isKey = p.immersion[len(p.immersion)-1] == imObject
		default:
			if p.keyAction == SkipAction && p.controlKeyDepth < len(p.immersion) {
				continue
			}
			if p.isKey {
				p.keyAction = keyHandler(fmt.Sprint(token), len(p.immersion)-1)
				if p.keyAction != ReadAction {
					p.controlKeyDepth = len(p.immersion)
				}
				continue
			}
			if p.keyAction == ReturnAction {

			}
			p.isKey = !p.isKey
		}
	}
	return nil
}
