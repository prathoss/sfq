package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Prathoss/sfq/parsers"
	"github.com/spf13/cobra"
)

var structure *string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sfq",
	Short: "Query structured files",
	Long: `Tool for manipulation of structured files.

Support for getting data out of the file or changing the data.`,
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

func getParser(fileName string) (parsers.Parser, error) {
	if *structure == "" {
		parsedFileName := strings.Split(fileName, ".")
		length := len(parsedFileName)
		if length < 2 {
			return nil, fmt.Errorf("Could not get structure neither from structure flag nor from file")
		}
		*structure = parsedFileName[length-1]
	}

	return parsers.GetParser(*structure)
}
