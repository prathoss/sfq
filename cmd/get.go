package cmd

import (
	"fmt"

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

		parser, err := getParser(file)
		if err != nil {
			return err
		}

		return parser.Parse(file,
			func(key string) bool {
				return key == query
			},
			func(value string) {
				fmt.Println(value)
			},
		)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
