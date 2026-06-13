package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/router"
	"github.com/gettako/tako/internal/tako"
)

// Layout defines the base Bubble Tea view structure.
type Layout struct{}

// View is the main render loop called by the Tako Bubble Tea adapter.
func (l *Layout) View(ctx *tako.Context, r *router.Router) tea.View {
	var overlayMgr contracts.OverlayManager
	_ = ctx.Container().Make(&overlayMgr)
	
	baseContent := ""
	if base := ctx.Hooks().Get("tako.overlay.overlay-base"); base != nil {
		baseContent = base.(string)
	}

	if overlayMgr != nil && overlayMgr.IsActive() && overlayMgr.Top() != "overlay-base" {
		topID := overlayMgr.Top()
		
		content := ctx.Hooks().Get("tako.overlay." + topID)
		if content != nil {
			// Stack them vertically for the demo
			combined := baseContent + "\n\n" + content.(string)
			v := tea.NewView(combined)
			v.AltScreen = true
			return v
		}
	}

	if baseContent != "" {
		v := tea.NewView(baseContent)
		v.AltScreen = true
		return v
	}

	// Fallback empty view
	v := tea.NewView("Tako Overlays App booting...\nPress Ctrl+C to exit.")
	v.AltScreen = true
	return v
}
