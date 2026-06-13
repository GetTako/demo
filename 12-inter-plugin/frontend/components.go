package frontend

import (
	"context"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
)

type Dashboard struct {
	ctx           *tako.Context
	log           func() []string
	weatherResult string
}

func (d *Dashboard) ID() string { return "inter-plugin-dashboard" }

func (d *Dashboard) Render() any {
	b := strings.Builder{}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))
	b.WriteString(titleStyle.Render("=== Inter-Plugin Communication ===") + "\n\n")

	// EventBus section
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Render("1. EventBus (Async Pub/Sub)") + "\n")
	b.WriteString("Press 'p' to emit 'frontend:ping'. Backend will reply with 'backend:pong'.\n")

	logs := d.log()
	for _, l := range logs {
		b.WriteString("> " + l + "\n")
	}
	b.WriteString("\n")

	// RPC section
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Render("2. RPC (Sync Request/Response)") + "\n")
	b.WriteString("Press 'w' to call 'weather:get' endpoint from Backend.\n")

	if d.weatherResult != "" {
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Render("Weather: " + d.weatherResult))
	} else {
		b.WriteString("Waiting for RPC call...")
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
	z := keys.Zone(d.ID())

	z.Bind("p", func() {
		// Asynchronous Ping
		d.ctx.Emit("frontend:ping", "Hello from Frontend!")
	})

	z.Bind("w", func() {
		d.weatherResult = "fetching..."

		// Synchronous RPC call to Backend Plugin
		res, err := d.ctx.RPC().Call("weather:get").
			WithPayload("Jakarta").
			WithContext(context.Background()).
			Await()

		if err != nil {
			d.weatherResult = "Error: " + err.Error()
		} else {
			d.weatherResult = res.Data.(string)
		}
	})
}
