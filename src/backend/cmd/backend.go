package cmd

import (
	"github.com/anggakes/gosvelte/src/backend/pkg/handlers"
	"github.com/anggakes/gosvelte/src/backend/pkg/providers"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "backend",
	Short: "Run backend service",
	Long:  `run backend service`,
	Run: func(cmd *cobra.Command, args []string) {

		e := echo.New()

		handlers.RegisterRoute(e, providers.InitProvider())

		e.Start(":8181")
	},
}
