package main

import (
	"github.com/bamboo-services/bamboo-base-go-template/internal/app/route"
	"github.com/bamboo-services/bamboo-base-go-template/internal/app/startup"
	xLog "github.com/bamboo-services/bamboo-base-go/log"
	xMain "github.com/bamboo-services/bamboo-base-go/main"
	xReg "github.com/bamboo-services/bamboo-base-go/register"
)

func main() {
	reg := xReg.Register(startup.Init())
	log := xLog.WithName(xLog.NamedMAIN)

	xMain.Runner(reg, log, route.NewRoute, nil)
	return
}
