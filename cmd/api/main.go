package main

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

func main() {
	Init()
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.Use(AccessLog())
	hlog.SetLevel(hlog.LevelTrace)
	rg := h.Group("/douyin")

	rg.GET("/feed/")

	// 用户登录接口
	rg.POST("/user/login/", handlers.UserAuthMw.LoginHandler)
	// 用户注册接口
	rg.POST("/user/register/", handlers.Register)
	// 用户信息
	rg.GET("/user/", handlers.UserAuthMw.MiddlewareFunc(), handlers.UserInfo)

	h.Spin()
}

func Init() {
	rpc.InitApiRpc()
	//
	handlers.Init()
}

func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Microseconds
		hlog.CtxTracef(c, "status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s query=%s",
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(),
			ctx.Request.Host(), ctx.Request.QueryString())
	}
}
