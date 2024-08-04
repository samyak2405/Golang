package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "Zero is a CLI tool for performing basic mathematical Operations",
	Long:  "Zero is a CLI tool for performing basic mathematical Operations - Addition, Multiplication, Subtraction and Division",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, "Oops, An error while Executing Zero %s\n", err)
		os.Exit(1)
	}
}
