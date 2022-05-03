package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Get grader version",
	Long:    "Get local grader version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("beta")
		fmt.Println("unstable project")
		fmt.Println("v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
