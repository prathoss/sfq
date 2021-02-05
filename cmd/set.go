package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Writes updated file by given query into standard output",
	Long:  `Writes updated file by given query into standard output.`,
	Args:  cobra.ExactValidArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
