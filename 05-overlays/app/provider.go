package app

import (

	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// OverlaysProvider is a ServiceProvider for the overlays demo.
type OverlaysProvider struct{}

func (p *OverlaysProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *OverlaysProvider) Boot(app *foundation.Application) error {
	app.Stack().Push("base")

	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		base := &BaseComponent{overlay: overlayMgr}
		overlayMgr.ShowComponent(base)
	}
	return nil
}


