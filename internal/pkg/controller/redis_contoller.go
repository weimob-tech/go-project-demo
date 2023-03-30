package controller

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wcontext"
)

func RedisGetController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		key := ctx.Param("key")
		foo, err := wcontext.Global().Redis().Get(c, key).Result()
		if err != nil {
			ctx.JSON(consts.StatusOK, x.Fail("90500", err.Error()))
			return
		}
		ctx.JSON(consts.StatusOK, x.OkAnd(foo))
	}
}

func RedisSetController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		type RedisSetReq struct {
			Key    string `json:"key"`
			Val    string `json:"val"`
			Expire int    `json:"expire"`
		}
		var req RedisSetReq
		err := ctx.Bind(&req)

		var hours = time.Duration(req.Expire)
		foo, err := wcontext.Global().Redis().Set(c, req.Key, req.Val, hours).Result()
		if err != nil {
			ctx.JSON(consts.StatusOK, x.Fail("90500", err.Error()))
			return
		}
		ctx.JSON(consts.StatusOK, x.OkAnd(foo))
	}
}
