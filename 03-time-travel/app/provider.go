package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/tako"
	"github.com/gettako/tako/pkg/foundation"
)

// TimeTravelProvider is a ServiceProvider for the time-travel demo.
type TimeTravelProvider struct{}

func (p *TimeTravelProvider) Register(_ *foundation.Application) error {
	return nil
}

func (p *TimeTravelProvider) Boot(app *foundation.Application) error {
	stateMgr := app.Context().State()

	bankComponent := &BankingComponent{
		ctx:     app.Context(),
		balance: 1000,
		state:   stateMgr,
	}

	stateMgr.Mutate("balance").Value(bankComponent.balance).Broadcast()

	stateMgr.Watch("balance").OnUpdate(func(oldVal, newVal any) {
		if val, ok := newVal.(int); ok {
			bankComponent.balance = val
		}
	}).Subscribe(app.Context().Context)

	app.UI().MountView(bankComponent)
	return nil
}

// ─── UI Component ─────────────────────────────────────────────────────────────

type BankingComponent struct {
	ctx     *tako.Context
	balance int
	state   contracts.StateManager
}

func (b *BankingComponent) ID() string { return "time-travel" }

func (b *BankingComponent) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))
	balanceStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#C7775D"))
	if b.balance < 0 {
		balanceStyle = balanceStyle.Foreground(lipgloss.Color("#C7775D"))
	}

	var sb strings.Builder
	sb.WriteString(titleStyle.Render("=== Tako Banking Simulation ===") + "\n\n")
	sb.WriteString(fmt.Sprintf("Current Balance: %s\n\n", balanceStyle.Render(fmt.Sprintf("$%d", b.balance))))
	sb.WriteString("Controls:\n")
	sb.WriteString("  [d] Deposit $100\n")
	sb.WriteString("  [w] Withdraw $50\n")
	sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Render("  [c] Simulate System Crash (Panic)\n"))
	sb.WriteString("  [ctrl+c] Graceful Quit\n")

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

func (b *BankingComponent) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(b.ID())

	z.Bind("d", func() { panic("Test crash boundary!")
		b.state.Mutate("balance").Value(b.balance + 100).Broadcast()
	})

	z.Bind("w", func() {
		b.state.Mutate("balance").Value(b.balance - 50).Broadcast()
	})

	z.Bind("c", func() {
		panic("CRITICAL ERROR: Banking system crashed! State corrupted!")
	})
}
