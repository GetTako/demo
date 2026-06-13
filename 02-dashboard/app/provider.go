package app

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
	"github.com/gettako/tako/pkg/foundation"
)

// DashboardProvider is a ServiceProvider for the Dashboard feature.
type DashboardProvider struct{}

func (p *DashboardProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *DashboardProvider) Boot(app *foundation.Application) error {
	dashboard := &DashboardBox{ctx: app.Context()}

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
				bus.Emit("sys:update", map[string]int{
					"cpu":    rand.Intn(100),
					"memory": rand.Intn(16384),
				})
			}
		}
	}()

	app.UI().MountLayout(&MainLayout{ui: app.UI()})
	app.UI().MountView(dashboard)
	return nil
}

// ─── Layout Component ─────────────────────────────────────────────────────────

type MainLayout struct {
	ui contracts.UIManager
}

func (l *MainLayout) ID() string { return "main-layout" }

func (l *MainLayout) Render() any {
	// 1. Sidebar
	sidebarStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#C7775D")).
		Width(30).
		Height(15).
		Padding(1)

	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D")).Render("TAKO MENU")
	sidebarContent := title + "\n\n" +
		"1. Dashboard\n" +
		"2. Settings\n" +
		"3. Logout"

	sidebar := sidebarStyle.Render(sidebarContent)

	// 2. Slot / View Content
	var slotContent string
	if view := l.ui.RenderView(); view != nil {
		if s, ok := view.(string); ok {
			slotContent = s
		}
	}

	slotStyle := lipgloss.NewStyle().MarginLeft(2)

	// 3. Combine them horizontally
	combined := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, slotStyle.Render(slotContent))
	
	// Add margin so it aligns with other demos that have Margin(2, 4)
	return lipgloss.NewStyle().Margin(2, 4).Render(combined)
}

func (l *MainLayout) RegisterKeys(keys contracts.KeyManager) {
	// Layout can define global keys since it is the outer shell
	keys.Bind("1", func() { /* Navigate to Dashboard */ })
	keys.Bind("2", func() { /* Navigate to Settings */ })
}

// ─── UI Component (View) ──────────────────────────────────────────────────────

type DashboardBox struct {
	ctx    *tako.Context
	cpu    int
	memory int
}

func (d *DashboardBox) ID() string { return "dashboard" }

func (d *DashboardBox) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))
	metricStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Bold(true)

	var b strings.Builder
	b.WriteString(titleStyle.Render("=== System Metrics ===") + "\n\n")

	cpuBarLength := d.cpu / 5
	cpuBar := strings.Repeat("█", cpuBarLength) + strings.Repeat("░", 20-cpuBarLength)

	b.WriteString(fmt.Sprintf("CPU Load: [%s] %s%%\n\n", metricStyle.Render(cpuBar), metricStyle.Render(fmt.Sprint(d.cpu))))
	b.WriteString(fmt.Sprintf("Memory  : %s MB\n", metricStyle.Render(fmt.Sprint(d.memory))))

	var termWidth int
	if d.ctx != nil {
		_ = d.ctx.Storage().Get("term_width", &termWidth)
	}
	if termWidth <= 0 {
		termWidth = 80
	}
	containerWidth := termWidth - 40
	if containerWidth < 40 {
		containerWidth = 40
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#C7775D")).
		Padding(1, 4).
		Width(containerWidth).
		Height(15)

	return containerStyle.Render(b.String())
}

func (d *DashboardBox) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(d.ID())
	z.Bind("q", func() {})
}
