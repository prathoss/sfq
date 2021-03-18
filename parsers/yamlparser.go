package parsers

import (
	"fmt"
	"io"
)

type yamlParser struct {
}

func (p *yamlParser) Parse(reader io.ReadSeekCloser, keyHandler func(string, int) KeyAction, valueHandler func(string), otherSymbolHandler func(rune)) error {
	return fmt.Errorf("Not implemented yet")
}
