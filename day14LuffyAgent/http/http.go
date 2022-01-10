package http

import (
	"day14LuffyAgent/logger"
	"day14LuffyAgent/settings"
	"github.com/valyala/fasthttp"
)

func Start()  {
	addr := settings.Config().Http.Listen
	if addr != ""{
		conn := func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/v1/check/health":
				agentStatusRoutes(ctx)
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		}
		conn = fasthttp.CompressHandler(conn)
		fasthttp.ListenAndServe(addr, conn)
	}else{
		logger.StartupInfo("cfg.conf中http的address配置有误")
	}
}
