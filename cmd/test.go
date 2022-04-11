package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "testing testing",
	Long:  "testing testing testing",
	Run: func(cmd *cobra.Command, args []string) {
		mydir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(mydir)

		goExc, err := exec.LookPath("go")
		if err != nil {
			fmt.Println(err.Error())
		}

		cmdGo := &exec.Cmd{
			Path:   goExc,
			Args:   []string{goExc, "get", "-u", "gorm.io/gorm"},
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}

		fmt.Println(cmdGo.String())

		if err := cmdGo.Run(); err != nil {
			fmt.Println("Error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
