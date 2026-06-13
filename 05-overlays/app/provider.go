package app

import (
	"github.com/gettako/tako/pkg/foundation"
)

// OverlaysProvider is a ServiceProvider for the overlays demo.
type OverlaysProvider struct{}

func (p *OverlaysProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *OverlaysProvider) Boot(app *foundation.Application) error {
	base := &BaseComponent{ctx: app.Context(), overlay: app.UI()}
	app.UI().MountView(base)
	return nil
}
