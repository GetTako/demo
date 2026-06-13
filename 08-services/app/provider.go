package app

import (
	"context"
	"errors"
	"time"

	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

type ServicesProvider struct {}

func (p *ServicesProvider) Register(_ *foundation.Application) error { return nil }

func (p *ServicesProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	// 1. Periodic Scheduler
	ctx.Jobs().Every(1*time.Second, func() {
		val, _ := ctx.State().Get("counter").(int)
		ctx.State().Set("counter", val+1)
	})

	// 2. Dispatch Error Task when user requests it via event
	ctx.Events().On("trigger:error", func(e contracts.Event) {
		ctx.Jobs().Dispatch(func(jobCtx context.Context) (any, error) {
			time.Sleep(1 * time.Second) // simulate work
			return nil, errors.New("simulated API failure")
		}).OnError(func(err error) {
			ctx.State().Set("last_error", err.Error())
		})
	})

	app.Stack().Push("base")
	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(&Dashboard{ctx: ctx})
	}
	return nil
}
