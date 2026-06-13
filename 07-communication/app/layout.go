package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/router"
	"github.com/gettako/tako/internal/tako"
)

// MainLayout defines the base Bubble Tea view structure.
type MainLayout struct{}

// View is the main render loop called by the Tako Bubble Tea adapter.
func (l *MainLayout) View(ctx *tako.Context, r *router.Router) tea.View {
	var overlayMgr contracts.OverlayManager
	_ = ctx.Container().Make(&overlayMgr)
	
	if overlayMgr != nil && overlayMgr.IsActive() {
		topID := overlayMgr.Top()
		content := ctx.Hooks().Get("tako.overlay." + topID)
		if content != nil {
			v := tea.NewView(content.(string))
			v.AltScreen = true
			return v
		}
	}

	// Fallback empty view
	v := tea.NewView("Communication Demo booting...\nPress Ctrl+C to exit.")
	v.AltScreen = true
	return v
}
