package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Prathoss/sfq/parsers"
	"github.com/spf13/cobra"
)

var (
	structure *string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sfq",
	Short: "Query structured files",
	Long: `Tool for manipulation of structured files.

Support for getting data out of the file or changing the data.
Expects valid file, when file is invalid may not behave correctly.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	structure = rootCmd.PersistentFlags().StringP("structure", "s", "", "set file structure if file extension does not match files structure")
}

//first argument is always query, second is optional (fileName)
var (
	argsSetup = cobra.RangeArgs(1, 2)
)

/*getSourceAndParser returns

- reader: either standard input if file not set

- func: func to close file

- parser: correct parser for file structure

- error
*/
func getSourceAndParser(args []string) (io.ReadSeekCloser, func(), parsers.Parser, error) {
	if len(args) == 1 {
		if *structure == "" {
			return nil, nil, nil, fmt.Errorf("structure flag has to be set when reading from pipe")
		}
		parser, err := parsers.GetParser(*structure)
		if err != nil {
			return nil, nil, nil, err
		}
		return os.Stdin, func() {}, parser, nil
	}

	fileName := args[1]
	fileReader, err := os.Open(fileName)
	if err != nil {
		return nil, nil, nil, err
	}
	parsedFileName := strings.Split(fileName, ".")
	fnLen := len(parsedFileName)
	if fnLen < 2 {
		return nil, nil, nil, fmt.Errorf("could not get structure from file extension")
	}
	str := parsedFileName[fnLen-1]
	parser, err := parsers.GetParser(str)
	if err != nil {
		return nil, nil, nil, err
	}
	return fileReader, func() { fileReader.Close() }, parser, nil
}
