package frontend

import (
	tea "charm.land/bubbletea/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/router"
	"github.com/gettako/tako/internal/tako"
)

type MainLayout struct{}

func (l *MainLayout) View(ctx *tako.Context, r *router.Router) tea.View {
	var overlayMgr contracts.OverlayManager
	_ = ctx.Container().Make(&overlayMgr)
	
	if overlayMgr != nil && overlayMgr.IsActive() {
		topID := overlayMgr.Top()
		if content := ctx.Hooks().Get("tako.overlay." + topID); content != nil {
			v := tea.NewView(content.(string))
			v.AltScreen = true
			return v
		}
	}

	v := tea.NewView("Loading Inter-Plugin Demo...")
	v.AltScreen = true
	return v
}
