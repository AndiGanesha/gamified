package application

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/AndiGanesha/gamified/configuration"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
)

const (
	AppName = "tinderKW"
)

type App struct {
	Name          string
	Configuration *configuration.Configuration
	Context       context.Context
	ContextCancel context.CancelFunc
	Redis         redis.Client
	DB            *sql.DB
}

func NewApp(ctx context.Context, ctxCancel context.CancelFunc) (*App, error) {
	// initiate app
	app := &App{
		Name:          AppName,
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

	// Connect to Redis
	address := appConfig.Redis.Host + ":" + strconv.Itoa(appConfig.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: appConfig.Redis.Password,
		DB:       0,
	})
	app.Redis = *client
	log.Println("connecting to Redis", address)

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 appConfig.DB.Username,
		Passwd:               appConfig.DB.Password,
		Net:                  appConfig.DB.Host,
		Addr:                 appConfig.DB.Port,
		DBName:               appConfig.DB.Name,
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connecting to DB", appConfig.DB.Port)

	app.DB = db

	return app, nil
}

func (app *App) Close() {
	// Close Redis Connection
	err := app.Redis.Close()
	if err != nil {
		log.Println("Error closing Redis connection:", err)
	}

	// Close DB Connection
	err = app.DB.Close()
	if err != nil {
		log.Println("Error closing DB connection:", err)
	}
}
