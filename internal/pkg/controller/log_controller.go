package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

// LogInfoController
// if you want print log with json format, please config log.debug=false
// if you want print log with pretty format, please config log.debug=true
func LogInfoController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wlog.I().Any("age", 30).Str("name", "jack").Msg("person info")
		ctx.JSON(consts.StatusOK, x.OkAnd("log info"))
	}
}

func LogDebugController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wlog.D().Any("age", 30).Str("name", "jack").Msg("person info")
		ctx.JSON(consts.StatusOK, x.OkAnd("log debug"))
	}
}

func LogWarnController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wlog.W().Any("age", 30).Str("name", "jack").Msg("person info")
		ctx.JSON(consts.StatusOK, x.OkAnd("log warn"))
	}
}

func LogErrorController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wlog.E().Any("age", 30).Str("name", "jack").Msg("person info")
		ctx.JSON(consts.StatusOK, x.OkAnd("log error"))
	}
}

func LogPanicController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wlog.P().Any("age", 30).Str("name", "jack").Msg("person info")
		ctx.JSON(consts.StatusOK, x.OkAnd("log fatal"))
	}
}

func LogFatalController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wlog.F().Any("age", 30).Str("name", "jack").Msg("person info")
		ctx.JSON(consts.StatusOK, x.OkAnd("log fatal"))
	}
}
