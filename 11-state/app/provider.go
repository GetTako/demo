package app

import (
	"github.com/gettako/tako/pkg/foundation"
)

type StateProvider struct{}

func (p *StateProvider) Register(_ *foundation.Application) error { return nil }

func (p *StateProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	app.UI().MountView(&Dashboard{ctx: ctx})
	return nil
}
