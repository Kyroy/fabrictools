package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/pkg/errors"
)

func init() {
	rootCmd.AddCommand(failCmd)
}

var failCmd = &cobra.Command{
	Use:   "fail",
	Short: "This is a fail command.",
	Long:  "This command will fail no matter what.",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("I will fail now!")
		return errors.New("because I can")
	},
}
