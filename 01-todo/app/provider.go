package app

import (
	"github.com/gettako/tako/pkg/foundation"
)

// TodoProvider is a ServiceProvider for the Todo feature.
type TodoProvider struct{}

// Register registers the provider.
func (p *TodoProvider) Register(_ *foundation.Application) error {
	return nil
}

// Boot boots the provider.
func (p *TodoProvider) Boot(app *foundation.Application) error {
	box := &TodoBox{
		ctx: app.Context(),
	}
	box.Init()

	app.UI().MountView(box)

	return nil
}
