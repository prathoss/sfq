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
	file      *string
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
	file = rootCmd.PersistentFlags().StringP("file", "f", "", "set file to read, if not set use standard input")
}

func getParser() (parsers.Parser, error) {
	if *structure == "" {
		if *file == "" {
			return nil, fmt.Errorf("file flag not set, please set structure flag to parse file correctly")
		}
		parsedFileName := strings.Split(*file, ".")
		length := len(parsedFileName)
		if length < 2 {
			return nil, fmt.Errorf("could not get structure neither from structure flag nor from file")
		}
		*structure = parsedFileName[length-1]
	}

	return parsers.GetParser(*structure)
}

func getSource() (io.Reader, error) {
	if *file == "" {
		return os.Stdin, nil
	}

	fileReader, err := os.Open(*file)
	if err != nil {
		return nil, err
	}
	return fileReader, nil
}
