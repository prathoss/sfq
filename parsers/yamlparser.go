package parsers

import "fmt"

type yamlParser struct {
}

func (p *yamlParser) Parse(fileName string, keyHandler func(string, int) bool, valueHandler func(string), otherSymbolHandler func(rune)) error {
	return fmt.Errorf("Not implemented yet")
}
