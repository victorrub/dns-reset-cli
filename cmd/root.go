package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dns-reset",
	Short: "DNS-Reset automates the process of setting up a new network location on Mac OS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Success!")
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
