package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var subcmd = &cobra.Command{
	Use:   "Substract",
	Short: "sub 2 number",
	Long:  "add 2 number to substeract them",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"adding 2 number %s and %s is %s \n", args[0], args[1], Substract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subcmd)
}
