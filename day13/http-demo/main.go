package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	hertz := server.New()
	hertz.POST("/sys", func(c context.Context, ctx *app.RequestContext) {
		ctx.Data(200, "text/plain;charset=utf-8", []byte("ok"))
	})
	hertz.Spin()
}
