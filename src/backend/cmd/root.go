package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "This is app to run",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}