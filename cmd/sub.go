package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var subcmd = &cobra.Command{
	Use:     "Substract",
	Aliases: []string{"sub"},
	Short:   "sub 2 number",
	Long:    "add 2 number to substeract them",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err, res := Substract(args[0], args[1])
		if err != nil {
			return err
		}

		fmt.Printf(
			"adding 2 number %s and %s is %s \n", args[0], args[1], res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(subcmd)
}
