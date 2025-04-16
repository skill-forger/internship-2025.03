package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the go-project command in Cobra Command structure
var rootCmd = &cobra.Command{
	Use:   "go-project",
	Short: "go-project root command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("golang project root command called")
	},
}

// Execute processes the root Command and call the Run function
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

// init loads environment variable into Viper
func init() {
	rootCmd.Flags().String("config", "./local.env", "config file")
	viper.SetConfigFile(rootCmd.Flag("config").Value.String())

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv()
}
