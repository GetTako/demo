package app

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// CommunicationProvider demonstrates EventBus, RPC, and Hooks.
type CommunicationProvider struct {
	chatLog []string
}

func (p *CommunicationProvider) Register(_ *foundation.Application) error {
	p.chatLog = []string{"System: Welcome to the Pub/Sub Chat!"}
	return nil
}

func (p *CommunicationProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	// 1. EVENT BUS (Pub/Sub)
	// We subscribe to the "chat:message" event.
	ctx.Events().On("chat:message", func(e contracts.Event) {
		msg := e.Data.(string)
		timestamp := time.Now().Format("15:04:05")
		p.chatLog = append(p.chatLog, fmt.Sprintf("[%s] %s", timestamp, msg))
		if len(p.chatLog) > 10 {
			p.chatLog = p.chatLog[1:]
		}
	})

	// 2. RPC (Request/Response)
	// We register a route that acts as a simple calculator.
	ctx.RPC().Route("math:add").Handle(func(_ context.Context, req contracts.RPCRequest) (contracts.RPCResponse, error) {
		time.Sleep(300 * time.Millisecond) // Simulate delay

		args, ok := req.Payload.([]int)
		if !ok || len(args) != 2 {
			return contracts.RPCResponse{}, fmt.Errorf("invalid payload: expected []int with 2 elements")
		}

		result := args[0] + args[1]
		return contracts.RPCResponse{Data: strconv.Itoa(result)}, nil
	})

	// 3. HOOKS (Extensibility)
	// We inject UI widgets into the "sidebar:widgets" hook slot.
	ctx.Hooks().Add("sidebar:widgets", func() any {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Border(lipgloss.NormalBorder())
		return style.Render("Widget 1: System Online")
	})

	ctx.Hooks().Add("sidebar:widgets", func() any {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#C7775D")).Border(lipgloss.NormalBorder())
		return style.Render("Widget 2: 0 Active Errors")
	})

	// Prepare dashboard component tree
	dashboard := &Dashboard{
		ctx: ctx,
		chat: &ChatBox{
			bus: ctx.Events(),
			log: func() []string { return p.chatLog },
		},
		rpc: &RpcWidget{
			rpc: ctx.RPC(),
		},
		hook: ctx.Hooks(),
	}

	app.UI().MountView(dashboard)

	return nil
}
