package cmd

import (
	"log"

	"golang-project/migrations/data"
	"golang-project/migrations/schema"

	"github.com/spf13/cobra"
)

// rollbackCmd represents the rollback command in Cobra Command structure
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "migrate the database schema for the go-project",
	Run:   runRollbackCmd,
}

// init adds the rollback command into the migration command
func init() {
	rollbackCmd.Flags().Bool("data", false, "migrate data flag")
	rollbackCmd.Flags().Bool("schema", false, "migrate schema flag")

	migrationCmd.AddCommand(rollbackCmd)
}

// runRollbackCmd initialize a new connection and rollback changes from the database
func runRollbackCmd(cmd *cobra.Command, args []string) {
	databaseConnection := newDatabaseConnection()
	databaseInstance, err := databaseConnection.Open()
	if err != nil {
		log.Fatal("database error:", err)
	}

	dataFlag := cmd.Flag("data")
	schemaFlag := cmd.Flag("schema")

	if dataFlag != nil && dataFlag.Value.String() == "true" {
		err = data.NewMigration(databaseInstance).RollbackLast()
		if err != nil {
			log.Fatal("data rollback error:", err)
		}
		log.Println("data rollback completed")
	}

	if schemaFlag != nil && schemaFlag.Value.String() == "true" {
		err = schema.NewMigration(databaseInstance).RollbackLast()
		if err != nil {
			log.Fatal("schema rollback error:", err)
		}
		log.Println("schema rollback completed")
	}

	log.Println("migration rollback finished")
}
