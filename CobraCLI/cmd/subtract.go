package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"subtraction"},
	Short:   "Subtract 2 numbers",
	Long:    "Carry out subtraction operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Subtraction of %s and %s = %s", args[0], args[1], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
