package app

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// DashboardProvider is a ServiceProvider for the Dashboard feature.
type DashboardProvider struct{}

func (p *DashboardProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *DashboardProvider) Boot(app *foundation.Application) error {
	dashboard := &DashboardBox{}

	bus := app.Events()
	bus.Subscribe(app.Context().Context, "sys:update", func(e contracts.Event) {
		data := e.Data.(map[string]int)
		dashboard.cpu = data["cpu"]
		dashboard.memory = data["memory"]
	})

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-app.Context().Context.Done():
				return
			case <-ticker.C:
				bus.Publish("sys:update", map[string]int{
					"cpu":    rand.Intn(100),
					"memory": rand.Intn(16384),
				})
			}
		}
	}()

	app.Stack().Push("base")

	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(dashboard)
	}
	return nil
}

// ─── UI Component ─────────────────────────────────────────────────────────────

type DashboardBox struct {
	cpu    int
	memory int
}

func (d *DashboardBox) ID() string { return "dashboard" }

func (d *DashboardBox) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	metricStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)

	var b strings.Builder
	b.WriteString(titleStyle.Render("=== Tako Dashboard ===") + "\n\n")

	cpuBarLength := d.cpu / 5
	cpuBar := strings.Repeat("█", cpuBarLength) + strings.Repeat("░", 20-cpuBarLength)

	b.WriteString(fmt.Sprintf("CPU Load: [%s] %s%%\n\n", metricStyle.Render(cpuBar), metricStyle.Render(fmt.Sprint(d.cpu))))
	b.WriteString(fmt.Sprintf("Memory  : %s MB\n", metricStyle.Render(fmt.Sprint(d.memory))))

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("39")).
		Padding(1, 4).
		Margin(2, 4)

	return containerStyle.Render(b.String())
}

func (d *DashboardBox) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(d.ID())
	z.Bind("q", func() {})
}
