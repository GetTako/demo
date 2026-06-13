package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/router"
)

// ─── UI Component ─────────────────────────────────────────────────────────────

type MouseComponent struct {
	clickX int
	clickY int
	router *router.Router
}

func (m *MouseComponent) ID() string { return "mouse-demo" }

func (m *MouseComponent) Render() any {
	m.router.Mouse().UpdateHitbox(1, "click-target", 5, 5, 40, 10)

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("220"))

	var b strings.Builder
	b.WriteString(titleStyle.Render("=== Tako Mouse Interactive Demo ===") + "\n\n")
	b.WriteString("I've registered a hitbox at (x:5, y:5, w:40, h:10)\n")
	b.WriteString("Try clicking your mouse around those coordinates!\n\n")

	if m.clickX != -1 {
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Render(
			fmt.Sprintf("You clicked inside the target at: X=%d, Y=%d", m.clickX, m.clickY),
		))
	} else {
		b.WriteString("Waiting for click...")
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("220")).
		Padding(1, 4).
		Margin(2, 4)

	return containerStyle.Render(b.String())
}

func (m *MouseComponent) RegisterKeys(keys contracts.KeyManager) {
}

func (m *MouseComponent) RegisterMouse(mouse contracts.MouseManager) {
	z := mouse.Zone("click-target")
	z.OnClick(func(x, y int) {
		m.clickX = x
		m.clickY = y
	})
}
