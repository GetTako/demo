package app

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
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
		balance: 1000,
		state:   stateMgr,
	}

	stateMgr.Key("balance").Value(bankComponent.balance).Broadcast()

	stateMgr.Observe("balance").OnUpdate(func(oldVal, newVal any) {
		if val, ok := newVal.(int); ok {
			bankComponent.balance = val
		}
	}).Subscribe(app.Context().Context)

	app.Stack().Push("base")

	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(bankComponent)
	}
	return nil
}

// ─── UI Component ─────────────────────────────────────────────────────────────

type BankingComponent struct {
	balance int
	state   contracts.StateManager
}

func (b *BankingComponent) ID() string { return "time-travel" }

func (b *BankingComponent) Render() any {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	balanceStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("82"))
	if b.balance < 0 {
		balanceStyle = balanceStyle.Foreground(lipgloss.Color("196"))
	}

	var sb strings.Builder
	sb.WriteString(titleStyle.Render("=== Tako Banking Simulation ===") + "\n\n")
	sb.WriteString(fmt.Sprintf("Current Balance: %s\n\n", balanceStyle.Render(fmt.Sprintf("$%d", b.balance))))
	sb.WriteString("Controls:\n")
	sb.WriteString("  [d] Deposit $100\n")
	sb.WriteString("  [w] Withdraw $50\n")
	sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("  [c] Simulate System Crash (Panic)\n"))
	sb.WriteString("  [ctrl+c] Graceful Quit\n")

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(1, 4).
		Margin(2, 4)

	return containerStyle.Render(sb.String())
}

func (b *BankingComponent) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(b.ID())

	z.Bind("d", func() {
		b.state.Key("balance").Value(b.balance + 100).Broadcast()
	})

	z.Bind("w", func() {
		b.state.Key("balance").Value(b.balance - 50).Broadcast()
	})

	z.Bind("c", func() {
		panic("CRITICAL ERROR: Banking system crashed! State corrupted!")
	})
}
