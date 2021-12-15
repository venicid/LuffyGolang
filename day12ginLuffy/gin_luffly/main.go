package main

import (
	"context"
	"fmt"
	"gin_luffly/async"
	views "gin_luffly/gallery/api"
	usersDB "gin_luffly/gallery/models"
	"gin_luffly/logging"
	"gin_luffly/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"go.uber.org/fx"
	"net/http"
	"os"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()
	app.Name = "lufflyweb"
	app.Usage = "Gin rest demo"
	app.Version = "0.0.0"
}

func loadConfig() (*settings.Config, error) {
	return settings.Load()
}

func newServer(lc fx.Lifecycle, cfg *settings.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), cors.Default())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler: r,
	}
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logging.DefaultLogger().Infof("Start to rest api server: %d", cfg.ServerConfig.Port)
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logging.DefaultLogger().Infof("Stopped rest api server")
				return srv.Shutdown(ctx)
			},
		})
	return r
}

func printAppInfo(cfg *settings.Config) {
	logging.DefaultLogger().Infow("app info", "config", cfg)
}

func runApplication() {
	// setup app + run server
	app := fx.New(
		fx.Provide(
			loadConfig,
			settings.NewDatabase,
			usersDB.NewUsersDB,
			views.NewHandler,
			// gin server
			newServer,
		),
		fx.Invoke(
			views.RouteV1,
			printAppInfo,
		),
	)
	app.Run()
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "launch Gin Server By boyleGu",
			Action: func(c *cli.Context) error {
				runApplication()
				return nil
			},
		},
		{
			Name:  "worker",
			Usage: "luaunch machinery worker",
			Action: func(c *cli.Context) error {
				if err := async.Worker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)

	// runApplication()
}
