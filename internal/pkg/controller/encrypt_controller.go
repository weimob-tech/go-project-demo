package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/weimob-tech/go-project-base/pkg/auth"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-boot/pkg/encrypt"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

func EncryptController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := auth.GetCCToken(c, "qa", "", "")
		req := encrypt.NewBosEncryptRequest(token, "18612341234")
		res, err := encrypt.DoRequest(req)
		if err != nil {
			wlog.I().Msgf("encrypt error: %v", err)
		}
		wlog.I().Msgf("encrypt result: %v", codec.ToJson(res))

		req = encrypt.NewBosDecryptRequest(token, res.Data.Result)
		res, err = encrypt.DoRequest(req)
		if err != nil {
			wlog.I().Msgf("decrypt error: %v", err)
		}
		if res.Data.Result != "18612341234" {
			wlog.I().Msgf("decrypt failed, expect: 18612341234, got: %s", res.Data.Result)
		}
		wlog.I().Msgf("encrypt result: %v", codec.ToJson(res))
	}
}
