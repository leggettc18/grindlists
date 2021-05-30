/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/leggettc18/grindlists/api/api"
	"github.com/leggettc18/grindlists/api/app"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.New()
		if err != nil {
			return err
		}
		api, err := api.New(app)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			app.ConsoleLogger.Info().Msg("signal caught. shutting down...")
			cancel()
		}()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer cancel()
			serveAPI(ctx, api)
		}()

		wg.Wait()
		return nil
	},
}

func serveAPI(ctx context.Context, api *api.API) {
	router := mux.NewRouter()
	api.Init(router)

	var server *http.Server
	var handler http.Handler

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	handler = cors(router)

	server = &http.Server{
		Addr: fmt.Sprintf(":%d", api.App.Config.Server.Port),
		Handler: handler,
		ReadTimeout: 2 * time.Minute,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			api.App.ConsoleLogger.Error().Err(err).Msg("Error occurred during server shutdown")
		}
		<-done
	}()

	api.App.ConsoleLogger.Info().Msgf("serving api at http://127.0.0.1:%d", api.App.Config.Server.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		api.App.ConsoleLogger.Error().Err(err).Msg("Error occured during server startup.")
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
