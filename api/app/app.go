package app

import (
	"os"

	"github.com/rs/zerolog"
)

type App struct {
	Config *Config
	Logger *zerolog.Logger
	ConsoleLogger *zerolog.Logger
}

func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	consoleLog := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
	app.ConsoleLogger = &consoleLog

	return app, nil
}