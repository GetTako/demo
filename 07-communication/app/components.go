package app

import (
	"context"
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/internal/hook"
)

// ─── Chat Component (Event Bus) ──────────────────────────────────────────────

type ChatBox struct {
	bus   contracts.EventBus
	input string
	log   func() []string
}

func (c *ChatBox) ID() string { return "comm-chatbox" }

func (c *ChatBox) Render() any {
	b := strings.Builder{}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42"))
	b.WriteString(titleStyle.Render("1. Pub/Sub Chat (EventBus)") + "\n\n")

	// Render log
	logs := c.log()
	logStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	for _, l := range logs {
		b.WriteString(logStyle.Render(l) + "\n")
	}

	b.WriteString("\n")
	inputStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	b.WriteString(inputStyle.Render("Say: " + c.input + "_"))

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Width(45).
		Height(15)

	return box.Render(b.String())
}

func (c *ChatBox) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(c.ID())
	
	chars := "abcdefghijklmnopqrstuvwxyz0123456789-!?,. "
	for _, ch := range chars {
		char := string(ch)
		z.Bind(char, func() { c.input += char })
	}
	z.Bind("backspace", func() {
		if len(c.input) > 0 {
			c.input = c.input[:len(c.input)-1]
		}
	})
	z.Bind("enter", func() {
		if strings.TrimSpace(c.input) != "" {
			c.bus.Publish("chat:message", c.input)
			c.input = ""
		}
	})
}

// ─── RPC Component ────────────────────────────────────────────────────────────

type RpcWidget struct {
	rpc    contracts.RPCBus
	count  int
	result string
}

func (r *RpcWidget) ID() string { return "comm-rpc" }

func (r *RpcWidget) Render() any {
	b := strings.Builder{}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("220"))
	b.WriteString(titleStyle.Render("2. Calculator (RPC)") + "\n\n")

	b.WriteString("Press 'c' to calculate: " + fmt.Sprintf("%d + 10", r.count) + "\n\n")
	
	if r.result != "" {
		resStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("82"))
		b.WriteString(resStyle.Render("Result from backend: " + r.result))
	} else {
		b.WriteString("Waiting for request...")
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("214")).
		Padding(1, 2).
		Width(45).
		Height(8).
		MarginTop(1)

	return box.Render(b.String())
}

func (r *RpcWidget) RegisterKeys(keys contracts.KeyManager) {
	z := keys.Zone(r.ID())
	z.Bind("c", func() {
		r.result = "calculating..."
		
		// Synchronous RPC call
		res, err := r.rpc.Call("math:add").
			WithPayload([]int{r.count, 10}).
			WithContext(context.Background()).
			Await()
			
		if err != nil {
			r.result = "Error: " + err.Error()
		} else {
			r.result = res.Data.(string)
			r.count += 10
		}
	})
}

// ─── Dashboard Component (Layout) ─────────────────────────────────────────────

type Dashboard struct {
	chat *ChatBox
	rpc  *RpcWidget
	hook hook.Registry
}

func (d *Dashboard) ID() string { return "comm-dashboard" }

func (d *Dashboard) Render() any {
	chatUI := d.chat.Render().(string)
	rpcUI := d.rpc.Render().(string)

	// 3. HOOKS: Render dynamically injected side-widgets
	hookUI := d.renderHooks()

	leftColumn := lipgloss.JoinVertical(lipgloss.Left, chatUI, rpcUI)
	layout := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, hookUI)

	return lipgloss.NewStyle().Margin(1, 2).Render(layout)
}

func (d *Dashboard) renderHooks() string {
	b := strings.Builder{}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	b.WriteString(titleStyle.Render("3. Dynamic Sidebar (Hooks)") + "\n\n")
	b.WriteString("These widgets are injected via hook slots!\n\n")

	widgets := d.hook.All("sidebar:widgets")
	if len(widgets) == 0 {
		b.WriteString("No widgets injected.")
	} else {
		for _, w := range widgets {
			if str, ok := w.(string); ok {
				b.WriteString(str + "\n\n")
			}
		}
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("135")).
		Padding(1, 2).
		Width(30).
		Height(26).
		MarginLeft(2)

	return box.Render(b.String())
}

func (d *Dashboard) RegisterKeys(keys contracts.KeyManager) {
	d.chat.RegisterKeys(keys)
	d.rpc.RegisterKeys(keys)
}
