package app

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
	"github.com/gettako/tako/pkg/foundation"
)

// DemoPlugin is a ServiceProvider that also exposes a Manifest (external plugin).
type DemoPlugin struct{}

func (p *DemoPlugin) Register(_ *foundation.Application) error {
	return nil
}

func (p *DemoPlugin) Boot(app *foundation.Application) error {
	component := &PluginComponent{
		ctx:     app.Context(),
		message: "Plugin successfully initialized and wired!",
	}

	app.UI().MountView(component)
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
	ctx     *tako.Context
	message string
}

func (p *PluginComponent) ID() string { return "plugin-ui" }

func (p *PluginComponent) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))

	var b strings.Builder
	b.WriteString(titleStyle.Render("=== Tako Plugins Demo ===") + "\n\n")
	b.WriteString(p.message + "\n\n")
	b.WriteString("This demo shows how to construct a ServiceProvider\n")
	b.WriteString("with HasManifest interface for plugin:list visibility.")

	var termWidth int
	if p.ctx != nil {
		_ = p.ctx.Storage().Get("term_width", &termWidth)
	}
	if termWidth <= 0 {
		termWidth = 80
	}
	containerWidth := termWidth - 8
	if containerWidth < 40 {
		containerWidth = 40
	}

	return lipgloss.NewStyle().Margin(2, 4).Border(lipgloss.DoubleBorder()).Padding(1, 2).Width(containerWidth).Render(b.String())
}

func (p *PluginComponent) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(p.ID())
	z.Bind("q", func() {})
}
