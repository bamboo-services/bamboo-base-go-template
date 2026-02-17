package startup

import (
	"context"

	"github.com/bamboo-services/bamboo-base-go-template/internal/app/startup/prepare"
	xLog "github.com/bamboo-services/bamboo-base-go/log"
)

func (r *reg) businessDataPrepare(ctx context.Context) (any, error) {
	log := xLog.WithName(xLog.NamedINIT)
	prepare.New(log, ctx).Prepare()
	return nil, nil
}
