package main

import (
	"context"
	"day13ginLuffy/async"
	views "day13ginLuffy/gallery/api"
	usersDB "day13ginLuffy/gallery/models"
	"day13ginLuffy/logging"
	"day13ginLuffy/settings"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"go.uber.org/fx"
	"net/http"
	"os"
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
			views.RouteV3,
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

// 初始方式启动服务
//func main()  {
//	runApplication()
//}


// cli启动app、worker
var (
	app *cli.App
)

func init()  {
	app = cli.NewApp()
	app.Name = "luffyWbe"
	app.Usage = "Gin rest demo"
	app.Version = "1.0.0"
}

func main()  {
	app.Commands = []cli.Command{
		{
			Name: "server",
			Usage: "lanch Gin Server By alex",
			Action: func(c *cli.Context) error {
				runApplication()
				return nil
			},
		},
		{
			Name: "worker",
			Usage: "lanch machinery worker",
			Action: func(c *cli.Context) error {
				if err := async.Worker(); err!=nil{
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)

}


/*
访问

1.启动gin
E:\golang\HelloGolang\day13ginLuffy>go run main.go server

2. 启动worker
E:\golang\HelloGolang\day13ginLuffy>go run main.go worker


3. 访问异步接口
curl --location --request GET '127.0.0.1:8080/api/v3/users/longTask'

4. 查看redis结果，20秒后执行
DEBUG: 2021/12/16 00:42:20 redis.go:347 Received new message: {"UUID":"task_cbef1615-425a-4928-a723-24f3d44466c2","Name":"sum","RoutingKey":"machineryDemo","ETA":null,"GroupUUID":"","Gro
upTaskCount":0,"Args":[{"Name":"","Type":"[]int64","Value":[1,2,3,4,5,6]}],"Headers":{},"Priority":0,"Immutable":false,"RetryCount":0,"RetryTimeout":0,"OnSuccess":null,"OnError":null,"Ch
ordCallback":null,"BrokerMessageGroupId":"","SQSReceiptHandle":"","StopTaskDeletionOnError":false,"IgnoreWhenTaskNotRegistered":false}?[0m

?[0;91mERROR: 2021/12/16 00:42:40 workers.go:35  error handler sum errors ?[0m
*/