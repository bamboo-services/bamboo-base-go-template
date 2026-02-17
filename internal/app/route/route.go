package route

import (
	"context"

	xMiddle "github.com/bamboo-services/bamboo-base-go/middleware"
	xReg "github.com/bamboo-services/bamboo-base-go/register"
	xRoute "github.com/bamboo-services/bamboo-base-go/route"
	"github.com/gin-gonic/gin"
)

type route struct {
	engine  *gin.Engine
	context context.Context
}

func NewRoute(reg *xReg.Reg) {
	r := &route{
		engine:  reg.Serve,
		context: reg.Init.Ctx,
	}

	r.engine.NoMethod(xRoute.NoMethod)
	r.engine.NoRoute(xRoute.NoRoute)

	r.engine.Use(xMiddle.ResponseMiddleware)
	r.engine.Use(xMiddle.ReleaseAllCors)
	r.engine.Use(xMiddle.AllowOption)

	apiRouter := r.engine.Group("/api/v1")
	r.healthRouter(apiRouter)
}
