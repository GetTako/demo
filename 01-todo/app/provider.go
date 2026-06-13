package app

import (
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// TodoProvider is a ServiceProvider for the Todo feature.
type TodoProvider struct{}

func (p *TodoProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *TodoProvider) Boot(app *foundation.Application) error {
	box := &TodoBox{
		ctx: app.Context(),
	}
	box.Init()

	app.Stack().Push("base")

	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(box)
	}
	return nil
}
