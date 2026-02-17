package handler

import (
	"context"

	"github.com/bamboo-services/bamboo-base-go-template/internal/logic"
	xLog "github.com/bamboo-services/bamboo-base-go/log"
)

type service struct {
	healthLogic *logic.HealthLogic
}

type handler struct {
	name    string
	log     *xLog.LogNamedLogger
	service *service
}

type IHandler interface {
	~struct {
		name    string
		log     *xLog.LogNamedLogger
		service *service
	}
}

func NewHandler[T IHandler](ctx context.Context, handlerName string) *T {
	return &T{
		name: handlerName,
		log:  xLog.WithName(xLog.NamedCONT, handlerName),
		service: &service{
			healthLogic: logic.NewHealthLogic(ctx),
		},
	}
}

type HealthHandler handler
