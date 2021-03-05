package parsers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Parser interface {
	Parse(io.Reader, func(string, int) bool, func(string), func(rune)) error
}

type StructureNotRecognisedError struct {
	structure string
}

func (pe StructureNotRecognisedError) Error() string {
	return fmt.Sprintf("Could not find parser for structure: %v", pe.structure)
}

func GetParser(structure string) (p Parser, err error) {
	if structure == "json" {
		p = &jsonParser{depth: -1}
		return
	}
	if structure == "yaml" || structure == "yml" {
		p = &yamlParser{}
		return
	}

	err = StructureNotRecognisedError{structure: structure}
	return
}

func readFile(fileName string, parse func(string)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	for {
		buffer := make([]byte, 16*1024)

		length, err := fileReader.Read(buffer)
		buffer = buffer[:length]

		if length == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}

		text := string(buffer)
		parse(text)
	}
	return nil
}
