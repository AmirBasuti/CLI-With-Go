package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "use 2 numbers to add them",
	Long: `This command takes 2 numbers and adds them together.
For example:
CLI add 2 3`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"adding 2 number %s and %s is %s \n", args[0], args[1], Add(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
