package parsers

import (
	"fmt"
	"io"
)

//KeyAction defines which action to take for the key
type KeyAction int

const (
	//ReadAction continue reading value of key
	ReadAction KeyAction = iota
	//ReturnAction return the value of key
	ReturnAction
	//SkipAction skip reading value of the key
	SkipAction
)

type immersionType int

const (
	imObject immersionType = iota
	imArray
)

type Parser interface {
	Parse(io.ReadSeekCloser, func(string, int) KeyAction, func(string), func(rune)) error
}

type StructureNotRecognisedError struct {
	structure string
}

func (pe StructureNotRecognisedError) Error() string {
	return fmt.Sprintf("Could not find parser for structure: %v", pe.structure)
}

func GetParser(structure string) (Parser, error) {
	if structure == "json" {
		parser := NewJsonParser()
		return &parser, nil
	}
	if structure == "yaml" || structure == "yml" {
		return &yamlParser{}, nil
	}

	return nil, StructureNotRecognisedError{structure: structure}
}
