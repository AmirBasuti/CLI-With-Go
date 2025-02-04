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

	RunE: func(cmd *cobra.Command, args []string) error {
		err, res := Add(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1], res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
