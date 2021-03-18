package cmd

import (
	"fmt"
	"strings"

	"github.com/Prathoss/sfq/parsers"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Writes updated file by given query into standard output",
	Long:  `Writes updated file by given query into standard output.`,
	Args:  argsSetup,
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		eqSplit := strings.Split(query, "=")
		if len(eqSplit) < 2 {
			return fmt.Errorf("Query must contain '='")
		}

		keys := strings.Split(eqSplit[0], ".")
		valueToSet := strings.Join(eqSplit[1:], "=")

		nOfKeys := len(keys) - 1

		reader, closeFunc, parser, err := getSourceAndParser(args)
		if err != nil {
			return err
		}
		defer closeFunc()

		otherSymbolsBuilder := strings.Builder{}
		isLastKey := false

		return parser.Parse(reader,
			func(key string, depth int) parsers.KeyAction {
				fmt.Print(otherSymbolsBuilder.String())
				otherSymbolsBuilder.Reset()
				fmt.Print(key)
				if depth == nOfKeys && keys[depth] == key {
					fmt.Print(valueToSet)
					return parsers.SkipAction
				}
				if depth < nOfKeys && keys[depth] == key {
					return parsers.ReadAction
				}
				return parsers.ReturnAction
			},
			func(value string) {
				fmt.Print(otherSymbolsBuilder.String())
				otherSymbolsBuilder.Reset()

				if isLastKey == true {
					fmt.Print(valueToSet)
					isLastKey = false
					return
				}
				fmt.Println(value)
			},
			func(r rune) {
				otherSymbolsBuilder.WriteRune(r)
			},
		)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
