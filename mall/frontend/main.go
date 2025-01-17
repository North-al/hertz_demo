// Code generated by hertz generator.

package main

import (
	"context"
	"time"

	"github.com/North-al/hertz_demo/mall/frontend/biz/router"
	"github.com/North-al/hertz_demo/mall/frontend/conf"
	"github.com/North-al/hertz_demo/mall/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// init dal
	// dal.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")

	h.GET("/about", middleware.Auth, func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about.tmpl", nil)
	})

	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {

		data := utils.H{
			"Next": ctx.Query("next"),
		}

		ctx.HTML(consts.StatusOK, "sign-in.tmpl", data)
	})

	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up.tmpl", nil)
	})

	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, err := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte("secret"))
	h.Use(sessions.New("northal-shop", store))
	if err != nil {
		panic(err)
	}

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	h.Use(middleware.AuthMiddleware())
}
