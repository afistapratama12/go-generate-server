package cmd

import (
	"go-gen-server/config"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize new service",
	Long:  "Initialize service with install default package gin-gonic and GORM",
	Run: func(cmd *cobra.Command, args []string) {
		// get now directory, create directory if null
		// get name folder directory
		// - go mod init name sesuai nama folder <nama folder>
		// - go get -u github.com/gin-gonic/gin
		// - go get -u gorm.io/gorm
		// create config, connection with mysql (default)
		// add helper custom
		// db default using mysql
		// --db-mysql
		// --db-postgres / --db-postgresql / --db-pg
		// --db-sqlite

		// test test
		config.CreateOrOpenDir(config.RepoDir)
		config.CreateOrOpenFile("user", config.RepoDir)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
