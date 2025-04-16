package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang-project/internal/middleware"
	"golang-project/internal/registry"
	"golang-project/server"
	"golang-project/static"
)

// serveCmd represents the serve command in Cobra Command structure
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "initialize and start up the go-project server",
	Run:   runServeCmd,
}

// init adds the serve command into the root command
func init() {
	rootCmd.AddCommand(serveCmd)
}

// runServeCmd executes the core logic of the serve command
func runServeCmd(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	databaseConnection := newDatabaseConnection()
	databaseInstance, err := databaseConnection.Open()
	if err != nil {
		log.Fatal("database error:", err)
	}

	err = databaseConnection.Ping()
	if err != nil {
		log.Fatal("database error:", err)
	}

	handlerRegistries, err := registry.NewHandlerRegistries(databaseInstance)
	if err != nil {
		log.Fatal("registry error:", err)
	}

	serverConfigs := []server.ConfigProvider{
		func(e *echo.Echo) { e.Debug = true },
		func(e *echo.Echo) { e.HTTPErrorHandler = middleware.ErrorHandler },
		func(e *echo.Echo) {
			e.Use(
				middleware.Recover(),
				middleware.Timeout(),
				middleware.Correlation(),
				middleware.Authentication(handlerRegistries),
			)
		},
	}

	serverEngine := server.NewEngine(viper.GetString(static.EnvServerAddress), serverConfigs...)

	go func() {
		log.Println("golang server starts on environment:", viper.GetString(static.EnvServerEnv))

		err = serverEngine.Startup(handlerRegistries...)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error:", err)
		}
	}()

	<-c

	err = databaseConnection.Close()
	if err != nil {
		log.Println(err)
	}

	err = serverEngine.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("golang server gracefully shutdowns")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}
