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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		fmt.Println("with args:", args)
		fmt.Println("and with structure flag:", *structure)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
