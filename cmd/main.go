package main

import (
	"github.com/weimob-tech/go-project-boot/pkg/boot"
	"github.com/weimob-tech/go-project-demo/internal/app"
)

func main() {
	starter := boot.Starter(
		boot.WithHttpServer(),
		boot.ConfigureHttpServer(app.SetupHttpServer),
		boot.RunAfterSetup(app.SetupApplication),
	)

	starter.Start()
}
