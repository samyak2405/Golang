package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var multipleCmd = &cobra.Command{
	Use:     "multiple",
	Aliases: []string{"multiple"},
	Short:   "Multiple 2 numbers",
	Long:    "Carry out multiplication operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Multiplication of %s and %s = %s", args[0], args[1], Multiple(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(multipleCmd)
}
