package application

import (
	"context"
	"log"

	"github.com/AndiGanesha/authentication/configuration"
)

const (
	AppName = "tinderKW"
)

type App struct {
	Name          string
	Configuration *configuration.Configuration
	Context       context.Context
	ContextCancel context.CancelFunc
}

func NewApp(ctx context.Context, ctxCancel context.CancelFunc) (*App, error) {
	// initiate app
	app := &App{
		Name:          AppName,
		Configuration: &configuration.Configuration{},
		Context:       ctx,
		ContextCancel: ctxCancel,
	}

	// load config from env
	appConfig, err := configuration.LoadConfiguration()
	if err != nil {
		log.Println("load config error", err)
		return nil, err
	}
	app.Configuration = appConfig

	return app, nil
}

func (app *App) Close() {
	// Close Redis Connection
}
