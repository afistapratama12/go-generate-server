package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "-v"},
	Short:   "Get grader version",
	Long:    "Get local grader version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
