package main

import (
	"context"
	views "day12ginLuffy/gallery/api"
	usersDB "day12ginLuffy/gallery/models"
	"day12ginLuffy/logging"
	"day12ginLuffy/settings"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)


// 启动gin
func runApplication()  {
	// set app + run server
	app := fx.New(
		fx.Provide(
			loadConfig,
			settings.NewDatabase,
			usersDB.NewUsersDB,
			views.NewHandler,
			newServer,  // gin server
			),
		fx.Invoke(
			views.RouteV1,
			views.RouteV2,
			printAppInfo,
			),
		)
	app.Run()

}

// 加载配置文件
func loadConfig() (*settings.Config, error)  {
	return settings.Load()
}

// 启动gin server
func newServer(lc fx.Lifecycle, cfg *settings.Config) *gin.Engine  {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), cors.Default())

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.ServerConfig.Port),
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
		},
	)
	return r
}

// 打印app信息
func printAppInfo(cfg *settings.Config)  {
	logging.DefaultLogger().Infow("app info", "config", cfg)
}

func main()  {
	runApplication()
}