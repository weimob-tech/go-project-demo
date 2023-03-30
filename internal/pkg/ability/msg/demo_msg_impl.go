package msg

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-ability-sdk/pkg/wmsg/wos/weimob_shop"
	"github.com/weimob-tech/go-ability-sdk/pkg/wmsg/xinyun/coupon"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

func DemoMsgImplWos() {
	shop, err := weimob_shop.NewService()
	if err != nil {
		panic(err)
	}
	shop.OnMallOrderMallOrderMessageEvent(
		func(ctx context.Context, message *msg.WosOpenMessage[weimob_shop.MallOrderMallOrderMessageEvent]) (response x.Result, err error) {
			wlog.I().Msgf("got request: %v", codec.ToJson(message))
			event := message.GetMsgBody()
			wlog.I().Msgf("got request body: %v", codec.ToJson(event))
			return x.Ok(), nil
		})
}

func DemoMsgImplXinyun() {
	client, err := coupon.NewService()
	if err != nil {
		panic(err)
	}
	client.OnCcCouponCreateCouponEvent(
		func(ctx context.Context, message *msg.XinyunOpenMessage[coupon.CcCouponCreateCouponEvent]) (response x.Result, err error) {
			wlog.I().Msgf("got request: %v", codec.ToJson(message))
			event := message.GetMsgBody()
			wlog.I().Msgf("got request body: %v", codec.ToJson(event))
			return x.Ok(), nil
		})
}
