package main

import (
	"bufio"
	"flag"
	"fmt"
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

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileReader := bufio.NewReader(file)
}
