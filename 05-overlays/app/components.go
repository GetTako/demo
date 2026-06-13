package app

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
)

// ─── Base UI Component ────────────────────────────────────────────────────────

type BaseComponent struct {
	overlay contracts.UIManager
	result  string
}

func (b *BaseComponent) ID() string { return "overlays-base" }

func (b *BaseComponent) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))

	var sb strings.Builder
	sb.WriteString(titleStyle.Render("=== Tako Overlays Demo ===") + "\n\n")
	sb.WriteString("Press 'd' to open a native Confirm Dialog\n")
	sb.WriteString("Press 'p' to open a custom Modal Popup\n\n")

	if b.result != "" {
		sb.WriteString("Last Action: " + lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(b.result) + "\n")
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(1, 4).
		Margin(2, 4)

	return containerStyle.Render(sb.String())
}

func (b *BaseComponent) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(b.ID())

	z.Bind("d", func() {
		b.overlay.Dialog().Confirm("Are you sure you want to proceed?", func(yes bool) {
			if yes {
				b.result = "User clicked YES on confirm dialog"
			} else {
				b.result = "User clicked NO on confirm dialog"
			}
		})
	})

	z.Bind("p", func() {
		popup := &CustomPopup{overlay: b.overlay}
		b.overlay.MountOverlay(popup)
	})
}

// ─── Custom Popup Component ───────────────────────────────────────────────────

type CustomPopup struct {
	overlay contracts.UIManager
}

func (c *CustomPopup) ID() string { return "custom-popup" }

func (c *CustomPopup) Render() any {
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render("This is a custom popup overlay!\nNotice how it captures keyboard focus.\n\nPress 'esc' to close.")

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("82")).
		Background(lipgloss.Color("236")).
		Padding(2, 4).
		Render(text)

	return lipgloss.Place(80, 20, lipgloss.Center, lipgloss.Center, box)
}

func (c *CustomPopup) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(c.ID())
	z.Bind("esc", func() {
		c.overlay.Close()
	})
}
