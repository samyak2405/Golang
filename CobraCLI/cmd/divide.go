package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var divideCmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"division"},
	Short:   "Division of 2 numbers",
	Long:    "Carry out Division operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Division of %s and %s = %s", args[0], args[1], Divide(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)
}
