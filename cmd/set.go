package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Writes updated file by given query into standard output",
	Long:  `Writes updated file by given query into standard output.`,
	Args:  cobra.ExactValidArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		file := args[1]

		eqSplit := strings.Split(query, "=")
		if len(eqSplit) < 2 {
			return fmt.Errorf("Query must contain '='")
		}

		keys := strings.Split(eqSplit[0], ".")
		valueToSet := strings.Join(eqSplit[1:], "=")

		parser, err := getParser(file)
		if err != nil {
			return err
		}

		otherSymbolsBuilder := strings.Builder{}
		isLastKey := false

		return parser.Parse(file,
			func(key string, depth int) bool {
				fmt.Print(otherSymbolsBuilder.String())
				otherSymbolsBuilder.Reset()
				fmt.Print(key)
				if depth >= len(keys){
					return false
				}
				keyMatch := key == keys[depth]
				if keyMatch && depth == len(keys) - 1 {
					isLastKey = true
				}
				return keyMatch
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
