package frontend

import (
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// FrontendProvider simulates a UI plugin that consumes data from the backend.
type FrontendProvider struct {
	chatLog []string
}

func (p *FrontendProvider) Register(_ *foundation.Application) error {
	p.chatLog = []string{"Frontend initialized. Press P to Ping Backend."}
	return nil
}

func (p *FrontendProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	// Listen for backend responses
	ctx.Events().On("backend:pong", func(e contracts.Event) {
		p.chatLog = append(p.chatLog, e.Data.(string))
	})

	dashboard := &Dashboard{
		ctx: ctx,
		log: func() []string { return p.chatLog },
	}

	app.Stack().Push("base")
	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(dashboard)
	}

	return nil
}
