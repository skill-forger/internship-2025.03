package cmd

import (
	"log"

	"golang-project/migrations/data"
	"golang-project/migrations/schema"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command in Cobra Command structure
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "seed the database data for the go-project",
	Run:   runMigrateCmd,
}

// init adds the migrate command into the migration command
func init() {
	migrateCmd.Flags().Bool("data", false, "migrate data flag")
	migrateCmd.Flags().Bool("schema", false, "migrate schema flag")

	migrationCmd.AddCommand(migrateCmd)
}

// runMigrateCmd initialize a new connection and migrate changes into the database
func runMigrateCmd(cmd *cobra.Command, args []string) {
	databaseConnection := newDatabaseConnection()
	databaseInstance, err := databaseConnection.Open()
	if err != nil {
		log.Fatal("database error:", err)
	}

	dataFlag := cmd.Flag("data")
	schemaFlag := cmd.Flag("schema")

	if schemaFlag != nil && schemaFlag.Value.String() == "true" {
		err = schema.NewMigration(databaseInstance).Migrate()
		if err != nil {
			log.Fatal("schema migration error:", err)
		}
		log.Println("schema migration completed")
	}

	if dataFlag != nil && dataFlag.Value.String() == "true" {
		err = data.NewMigration(databaseInstance).Migrate()
		if err != nil {
			log.Fatal("data migration error:", err)
		}
		log.Println("data migration completed")
	}

	log.Println("migration migrate finished")
}
