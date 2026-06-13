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
}

func (d *Dashboard) ID() string { return "services-dashboard" }

func (d *Dashboard) Render() any {
	b := strings.Builder{}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))
	b.WriteString(titleStyle.Render("=== Services Demo (Scheduler) ===") + "\n\n")

	count, _ := d.ctx.State().Get("counter").(int)
	b.WriteString(fmt.Sprintf("Auto-incremented by Every(): %d\n\n", count))

	b.WriteString("Press 'e' to trigger a failing background job.\n")

	if errStr, ok := d.ctx.State().Get("last_error").(string); ok && errStr != "" {
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D"))
		b.WriteString(errStyle.Render("Last Background Error: " + errStr))
	}

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
		Margin(2, 4).
		Width(containerWidth)

	return box.Render(b.String())
}

func (d *Dashboard) RegisterKeys(keys contracts.KeyManager) {
	keys.Zone(d.ID()).Bind("e", func() {
		d.ctx.State().Set("last_error", "Processing...")
		d.ctx.Emit("trigger:error", nil)
	})
}
