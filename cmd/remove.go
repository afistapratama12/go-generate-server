package cmd

import (
	"fmt"
	"go-gen-server/internal"
	"os"

	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "remove entity service created",
	Long: `Remove entity service created
	
example : go-gen remove users
example2 : go-gen remove users,movies,todos
	
note : user '.' to remove all entity service in current directory
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running remove generate")

		if len(args) < 1 {
			fmt.Println("please insert entity name, ex : go-gen remove users")
			fmt.Println(`if you want to delete all entity service this directory please using "." , ex : go-gen remove .`)
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("error: cannot using the option flag, please try again")
			os.Exit(1)
		} else {
			ProcessRemove(args[0])
		}
	},
}

func ProcessRemove(nameEntity string) {
	internal.ProcessRemove(nameEntity)

}

func init() {
	rootCmd.AddCommand(remoteCmd)
}
