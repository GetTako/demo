package app

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
)

type Dashboard struct {
	ctx *tako.Context
}

func (d *Dashboard) ID() string { return "theming-dashboard" }

func (d *Dashboard) Render() any {
	th := d.ctx.Theme()
	la := d.ctx.Lang()

	b := strings.Builder{}

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(th.Get("accent")))
	b.WriteString(titleStyle.Render("=== Theming & i18n Demo ===") + "\n\n")

	b.WriteString("Greeting (Translation): " + la.T("greeting") + "\n\n")

	b.WriteString("Current Theme: " + th.Active() + "\n")
	b.WriteString("Current Lang: " + la.Active() + "\n\n")

	b.WriteString("Press 't' to toggle theme.\n")
	b.WriteString("Press 'l' to toggle language.\n")

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
		Background(lipgloss.Color(th.Get("bg"))).
		Foreground(lipgloss.Color(th.Get("fg"))).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(th.Get("accent"))).
		Padding(1, 2).
		Margin(2, 4).
		Width(containerWidth)

	return box.Render(b.String())
}

func (d *Dashboard) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(d.ID())

	z.Bind("t", func() {
		if d.ctx.Theme().Active() == "dark" {
			d.ctx.Theme().Use("light")
		} else {
			d.ctx.Theme().Use("dark")
		}
	})

	z.Bind("l", func() {
		if d.ctx.Lang().Active() == "en" {
			d.ctx.Lang().SetLocale("id")
		} else {
			d.ctx.Lang().SetLocale("en")
		}
	})
}
