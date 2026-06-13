package app

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// DemoPlugin is a ServiceProvider that also exposes a Manifest (external plugin).
type DemoPlugin struct{}

func (p *DemoPlugin) Register(_ *foundation.Application) error {
	return nil
}

func (p *DemoPlugin) Boot(app *foundation.Application) error {
	component := &PluginComponent{
		message: "Plugin successfully initialized and wired!",
	}

	app.Stack().Push("base")

	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(component)
	}
	return nil
}

// Manifest marks this provider as an external plugin visible in `plugin:list`.
func (p *DemoPlugin) Manifest() foundation.PluginManifest {
	return foundation.PluginManifest{
		Name:        "Demo Plugin",
		Version:     "1.0.0",
		Author:      "Tako Team",
		Description: "A demo plugin showing HasManifest interface",
	}
}

// ─── UI Component ─────────────────────────────────────────────────────────────

type PluginComponent struct {
	message string
}

func (p *PluginComponent) ID() string { return "plugin-ui" }

func (p *PluginComponent) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("87"))

	var b strings.Builder
	b.WriteString(titleStyle.Render("=== Tako Plugins Demo ===") + "\n\n")
	b.WriteString(p.message + "\n\n")
	b.WriteString("This demo shows how to construct a ServiceProvider\n")
	b.WriteString("with HasManifest interface for plugin:list visibility.")

	return lipgloss.NewStyle().Margin(2, 4).Border(lipgloss.DoubleBorder()).Padding(1, 2).Render(b.String())
}

func (p *PluginComponent) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(p.ID())
	z.Bind("q", func() {})
}
