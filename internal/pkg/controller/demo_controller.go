package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

// DemoGetController 展示最基础的 http controller 配置方式，包括获取参数、query、header、body 等
// 请求路径 http://localhost:8080/api/v1/demo/james/hello?foo=bar
func DemoGetController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		user := ctx.Param("user")
		foo := ctx.Query("foo")
		wlog.I().Str("query", foo).Str("user", user).Msg("got user info")
		ctx.JSON(consts.StatusOK, x.OkAnd("hello "+user))
	}
}

// DemoPostController post请求
func DemoPostController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		type personReq struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var req personReq
		_ = ctx.Bind(&req)

		ctx.JSON(consts.StatusOK, x.OkAnd(&req))
	}
}
