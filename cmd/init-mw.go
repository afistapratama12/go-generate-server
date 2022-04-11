package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initMwCmd = &cobra.Command{
	Use:     "init-middleware",
	Aliases: []string{"init-mw", "-init-mw"},
	Short:   "init-middleware to initialize middleware",
	Long:    "init-middleware to initialize middleware with auth function and middlware function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running init middleware")
	},
}

func init() {
	rootCmd.AddCommand(initMwCmd)
}
