package cmd

import (
	"fmt"
	"strings"

	"github.com/Prathoss/sfq/parsers"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get values from a file by given query",
	Long:  `Get values from a file by given query.`,
	Args:  argsSetup,
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		keys := strings.Split(query, ".")

		reader, closeFunc, parser, err := getSourceAndParser(args)
		if err != nil {
			return err
		}
		defer closeFunc()

		return parser.Parse(reader,
			func(key string, depth int) parsers.KeyAction {
				if depth == len(keys) - 1 && keys[depth] == key {
					return parsers.ReturnAction
				}
				if depth < len(keys) - 1 && keys[depth] == key {
					return parsers.ReadAction
				}
				return parsers.SkipAction
			},
			func(value string) {
				fmt.Println(value)
			},
			nil,
		)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
