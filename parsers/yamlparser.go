package parsers

import "fmt"

type yamlParser struct {
}

func (p *yamlParser) Parse(fileName string, keyHandler func(string) bool, valueHandler func(string)) error {
	return fmt.Errorf("Not implemented yet")
}
