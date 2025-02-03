package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var roundkonim bool
var multcmd = &cobra.Command{
	Use:   "Multiply",
	Short: "Multiply 2 number",
	Long:  "inter 2 number to Multiply them",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"adding 2 number %s and %s is %s \n", args[0], args[1], Multiply(args[0], args[1], roundkonim))
	},
}

func init() {
	multcmd.Flags().BoolVarP(&roundkonim, "round", "r", false, "round up 2 decimal number")
	rootCmd.AddCommand(multcmd)
}
