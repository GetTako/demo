package app

import (
	"github.com/gettako/tako/pkg/foundation"
)

// MouseProvider is a ServiceProvider for the mouse demo.
type MouseProvider struct{}

func (p *MouseProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *MouseProvider) Boot(app *foundation.Application) error {
	mouseDemo := &MouseComponent{
		ctx:    app.Context(),
		clickX: -1,
		clickY: -1,
		router: app.Router(),
	}

	app.UI().MountView(mouseDemo)
	return nil
}
