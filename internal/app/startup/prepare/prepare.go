package prepare

import (
	"context"

	xLog "github.com/bamboo-services/bamboo-base-go/log"
	xCtxUtil "github.com/bamboo-services/bamboo-base-go/utility/ctxutil"
	"gorm.io/gorm"
)

type Prepare struct {
	log *xLog.LogNamedLogger
	db  *gorm.DB
	ctx context.Context
}

func New(log *xLog.LogNamedLogger, ctx context.Context) *Prepare {
	return &Prepare{
		log: log,
		db:  xCtxUtil.MustGetDB(ctx),
		ctx: ctx,
	}
}

func (p *Prepare) Prepare() {
	p.prepareRole()
}
