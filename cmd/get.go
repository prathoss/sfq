package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get values from a file by given query",
	Long:  `Get values from a file by given query.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		file := args[1]

		keys := strings.Split(query, ".")

		parser, err := getParser(file)
		if err != nil {
			return err
		}

		return parser.Parse(file,
			func(key string, depth int) bool {
				if depth >= len(keys){
					return false
				}
				return key == keys[depth]
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
