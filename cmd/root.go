package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccheck",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
