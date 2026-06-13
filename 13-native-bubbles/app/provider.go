package app

import (
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// Provider is the demo provider.
type Provider struct{}

// Register registers the provider.
func (p *Provider) Register(_ *foundation.Application) error { return nil }

// Boot boots the provider.
func (p *Provider) Boot(app *foundation.Application) error {
	// Let's use foundation.Make to resolve the UIManager for demonstration.
	uiMgr, err := foundation.Make[contracts.UIManager](app)
	if err == nil {
		// Use the UI Manager to mount our Native Component.
		uiMgr.MountView(NewNativeInput())
	}
	return nil
}
