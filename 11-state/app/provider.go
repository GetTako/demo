package app

import (
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

type StateProvider struct {}

func (p *StateProvider) Register(_ *foundation.Application) error { return nil }

func (p *StateProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	app.Stack().Push("base")
	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(&Dashboard{ctx: ctx})
	}
	return nil
}
