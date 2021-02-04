package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	operation = flag.String("o", "get", "Operation on a file, available operation: get, set")
	query     = flag.String("q", "", "Query for selecting keys")
	fileName  = flag.String("f", "", "File to process")
	structure = flag.String("s", "json", "How is the file structured, available structures: json, yaml")
)

func main() {
	flag.Parse()
	if *operation != "get" && *operation != "set" {
		fmt.Println("Incorrect operation:", *operation)
		flag.Usage()
		return
	}

	if *structure != "json" && *structure != "yaml" {
		fmt.Println("Incorrect structure:", *structure)
		flag.Usage()
		return
	}
	var parser parser
	if *structure == "json" {
		parser = &jsonParser{symbol: false, key: false}
	} else {
		parser = &yamlParser{}
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		return
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
				fmt.Println(err.Error())
				os.Exit(1)
				return
			}
		}

		text := string(buffer)
		parser.parse(text)
	}
}

type parser interface {
	parse(string)
}

type jsonParser struct {
	symbol bool
	key    bool
}

func (parser *jsonParser) parse(text string) {
	for _, runeVal := range text {
		char := string(runeVal)
		if char == " " || char == "\n" {
			continue
		}
		if char == "\"" {
			parser.symbol = !parser.symbol
			fmt.Println()
			continue
		}
		fmt.Print(char)
	}
}

type yamlParser struct {
}

func (parser *yamlParser) parse(text string) {
	fmt.Println("Parsing yaml")
}
