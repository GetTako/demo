package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
)

type Dashboard struct {
	ctx *tako.Context

	// A fully decoupled component that doesn't handle inputs
	display *DisplayComponent
}

func (d *Dashboard) ID() string { return "state-dashboard" }

func (d *Dashboard) Render() any {
	if d.display == nil {
		d.display = &DisplayComponent{ctx: d.ctx}
		d.display.Init() // Subscribe to state
	}

	b := strings.Builder{}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))
	b.WriteString(titleStyle.Render("=== Reactive State Demo ===") + "\n\n")

	b.WriteString("Press 'UP' to increment state.\n")
	b.WriteString("Press 'DOWN' to decrement state.\n\n")

	// Render the inner component
	b.WriteString(d.display.Render().(string))

	var termWidth int
	_ = d.ctx.Storage().Get("term_width", &termWidth)
	if termWidth <= 0 {
		termWidth = 80
	}
	containerWidth := termWidth - 8
	if containerWidth < 40 {
		containerWidth = 40
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#C7775D")).
		Padding(1, 2).
		Width(containerWidth)

	return box.Render(b.String())
}

func (d *Dashboard) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(d.ID())
	z.Bind("up", func() {
		val, _ := d.ctx.State().Get("score").(int)
		d.ctx.State().Set("score", val+1)
	})
	z.Bind("down", func() {
		val, _ := d.ctx.State().Get("score").(int)
		d.ctx.State().Set("score", val-1)
	})
}

// ─── Fully Decoupled Component ────────────────────────────────────────────────

type DisplayComponent struct {
	ctx         *tako.Context
	latestScore int
}

func (dc *DisplayComponent) Init() {
	// Automatically reacts to state changes
	dc.ctx.State().Watch("score").OnUpdate(func(oldVal, newVal any) {
		if score, ok := newVal.(int); ok {
			dc.latestScore = score
		}
	}).Subscribe(dc.ctx.Context)
}

func (dc *DisplayComponent) Render() any {
	box := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 1)
	return box.Render(fmt.Sprintf("Current Score: %d", dc.latestScore))
}
