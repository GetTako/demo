package app

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
)

// ─── Base UI Component ────────────────────────────────────────────────────────

type BaseComponent struct {
	ctx     *tako.Context
	overlay contracts.UIManager
	result  string
}

func (b *BaseComponent) ID() string { return "overlays-base" }

func (b *BaseComponent) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))

	var sb strings.Builder
	sb.WriteString(titleStyle.Render("=== Tako Overlays Demo ===") + "\n\n")
	sb.WriteString("Press 'd' to open a native Confirm Dialog\n")
	sb.WriteString("Press 'p' to open a custom Modal Popup\n\n")

	if b.result != "" {
		sb.WriteString("Last Action: " + lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Render(b.result) + "\n")
	}

	var termWidth int
	if b.ctx != nil {
		_ = b.ctx.Storage().Get("term_width", &termWidth)
	}
	if termWidth <= 0 {
		termWidth = 80
	}
	containerWidth := termWidth - 8
	if containerWidth < 40 {
		containerWidth = 40
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("#C7775D")).
		Padding(1, 4).
		Margin(2, 4).
		Width(containerWidth)

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
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Render("This is a custom popup overlay!\nNotice how it captures keyboard focus.\n\nPress 'esc' to close.")

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#C7775D")).
		Background(lipgloss.Color("#C7775D")).
		Padding(2, 4).
		Render(text)

	return lipgloss.Place(80, 20, lipgloss.Center, lipgloss.Center, box)
}

func (c *CustomPopup) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(c.ID())
	z.Bind("esc", func() {
		c.overlay.Unmount()
	})
}
