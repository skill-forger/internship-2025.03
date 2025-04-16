package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang-project/database"
	"golang-project/static"
)

// migrationCmd represents the migration command in Cobra Command structure
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "go-project migration command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migration called")
	},
}

// init adds the migrate command into the root command
func init() {
	rootCmd.AddCommand(migrationCmd)
}

// newDatabaseConnection creates new database connection
func newDatabaseConnection() database.Connection {
	databaseSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString(static.EnvDbUser),
		viper.GetString(static.EnvDbPassword),
		viper.GetString(static.EnvDbHost),
		viper.GetString(static.EnvDbPort),
		viper.GetString(static.EnvDbName),
	)

	return database.NewConnection(databaseSourceName, nil)
}
