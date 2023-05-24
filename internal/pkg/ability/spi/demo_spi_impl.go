package spi

import (
	"context"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
	"github.com/weimob-tech/go-ability-sdk/pkg/wspi/wos/weimob_shop"
	"github.com/weimob-tech/go-ability-sdk/pkg/wspi/xinyun"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

func DemoSpiImplWos() {
	shop, err := weimob_shop.NewService()
	if err != nil {
		panic(err)
	}
	shop.OnPaasWeimobShopCouponPaasBatchLockCouponServiceInvocation(
		func(ctx context.Context, request *spi.WosInvocationRequest[weimob_shop.PaasWeimobShopCouponPaasBatchLockCouponRequest]) (response spi.InvocationResponse[weimob_shop.PaasWeimobShopCouponPaasBatchLockCouponResponse], err error) {
			wlog.I().Msgf("got request: %v", codec.ToJson(request))

			ret := weimob_shop.PaasWeimobShopCouponPaasBatchLockCouponResponse{Success: true}
			return spi.Ok(&ret), nil
		})
}

func DemoSpiImplXinyun() {
	shop, err := xinyun.NewService()
	if err != nil {
		panic(err)
	}
	shop.OnPaasAcquireCouponCodeServiceInvocation(
		func(ctx context.Context, request *spi.XinyunInvocationRequest[xinyun.PaasAcquireCouponCodeRequest]) (response spi.InvocationResponse[xinyun.PaasAcquireCouponCodeResponse], err error) {
			wlog.I().Msgf("got request: %v", codec.ToJson(request))

			ret := xinyun.PaasAcquireCouponCodeResponse{CouponList: []xinyun.PaasAcquireCouponCodeResponseCouponList{
				xinyun.PaasAcquireCouponCodeResponseCouponList{CouponTemplateId: 1, Code: "code", Wid: 34343, ErrorCode: 333, ErrMsg: "error"},
			}, HasCode: true}
			return spi.Ok(&ret), nil
		})
}
