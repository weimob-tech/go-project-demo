package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-ability-sdk/pkg/wapi/wos/weimob_shop"
	"github.com/weimob-tech/go-ability-sdk/pkg/wapi/xinyun/ec"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

func WeimobApiWosController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		clientInfo := auth.GetClientInfo("qa")
		client, err := weimob_shop.NewClientWithClientKey(clientInfo.ClientId, clientInfo.ClientSecret)
		if err != nil {
			panic(err)
		}

		req := weimob_shop.CreateCommentGoodsGetRequest()
		req.ShopId = "xxx"
		req.ShopType = "yyy"
		req.QueryParams = map[string]string{"accesstoken": auth.GetCCToken(c, "qa", "", "")}
		// req.QueryParameter = weimob_shop.CommentGoodsGetRequestQueryParameter{CommentIds: []int64{1, 2}}
		req.BasicInfo = weimob_shop.CommentGoodsGetRequestBasicInfo{Vid: 55}
		res, err := client.CommentGoodsGet(req)
		if err != nil {
			wlog.I().Msgf("api response: %v", err)
		}

		ctx.JSON(consts.StatusOK, x.OkAnd(res))
	}
}

func WeimobApiXinyunController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		client, err := ec.NewClientWithClientKey("", "")
		if err != nil {
			panic(err)
		}

		req := ec.CreateActivityQueryConsumeActivitySimpleResultRequest()
		req.QueryParams = map[string]string{"accesstoken": "34343434343"}
		req.Pid = 3434343
		req.OrderNo = "45454545"
		req.ActivityId = 434343434
		req.Wid = 45454545
		res, err := client.ActivityQueryConsumeActivitySimpleResult(req)
		if err != nil {
			wlog.I().Msgf("api response: %v", err)
		}

		ctx.JSON(consts.StatusOK, x.OkAnd(res))
	}
}
