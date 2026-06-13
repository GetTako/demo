package app

import (
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

type Provider struct{}

func (p *Provider) Register(_ *foundation.Application) error { return nil }

func (p *Provider) Boot(app *foundation.Application) error {
	// Let's use foundation.Make to resolve the UIManager for demonstration.
	uiMgr, err := foundation.Make[contracts.UIManager](app)
	if err == nil {
		// Use the UI Manager to mount our Native Component.
		uiMgr.MountView(NewNativeInput())
	}
	return nil
}
