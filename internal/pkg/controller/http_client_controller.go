package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

func HttpClientGetInvokeController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		cli := http.NewHttpClient().GetClient().(*client.Client)
		status, body, _ := cli.Get(context.Background(), nil, "https://www.baidu.com")
		wlog.I().Any("status", status).Str("body", string(body)).Msg("http client invoke get")
		ctx.JSON(consts.StatusOK, x.OkAnd("ok"))
	}
}

func HttpClientPostInvokeController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		req := &protocol.Request{}
		res := &protocol.Response{}
		req.SetMethod(consts.MethodPost)
		req.Header.SetContentTypeBytes([]byte("application/json"))
		req.Header.Set("X-Access-Code", "1234")
		reqMap := map[string]string{"projectId": "titan"}
		dst, _ := codec.Json.Marshal(reqMap)
		req.SetBodyRaw(dst)
		req.SetRequestURI("https://test.internal.weimobdev.com/testPath")
		cli := http.NewHttpClient().GetClient().(*client.Client)
		err := cli.Do(context.Background(), req, res)
		if err != nil {
			wlog.I().Str("body", string(res.Body())).Msg("http client invoke post error")
		}

		ctx.JSON(consts.StatusOK, x.OkAnd("ok"))
	}
}
