package app

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/msg"
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/spi"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-demo/internal/app/route"
	msgImpl "github.com/weimob-tech/go-project-demo/internal/pkg/ability/msg"
	spiImpl "github.com/weimob-tech/go-project-demo/internal/pkg/ability/spi"
)

func SetupHttpServer(s http.Server) {
	s.AddExtendCallback(msg.NewWosCallbackConfig())
	s.AddExtendCallback(msg.NewXinyunCallbackConfig())
	s.AddExtendCallback(spi.NewSpiCallbackConfig())
	route.Route(s.GetServer().(*server.Hertz))
}

func SetupApplication() {
	msgImpl.DemoMsgImplWos()
	msgImpl.DemoMsgImplXinyun()

	spiImpl.DemoSpiImplWos()
	spiImpl.DemoSpiImplXinyun()
}
