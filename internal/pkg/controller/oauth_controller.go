package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/auth"
)

func AccessTokenController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := auth.GetCCToken(c, "qa", "", "")
		ctx.String(consts.StatusOK, token)
	}
}
