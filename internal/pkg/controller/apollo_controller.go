package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/x"
)

func ApolloController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		val := config.GetString("my.client_id")
		ctx.JSON(consts.StatusOK, x.OkAnd(val))
	}
}
