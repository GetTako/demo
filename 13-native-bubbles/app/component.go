package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/gettako/tako/contracts"
)

// We won't use bubbles/textinput directly to avoid go.mod dependency issues,
// but we will simulate a raw BubbleTea model receiving events natively.
type NativeInput struct {
	text string
}

func NewNativeInput() *NativeInput {
	return &NativeInput{}
}

func (n *NativeInput) ID() string {
	return "native-input"
}

func (n *NativeInput) Render() any {
	var sb strings.Builder
	sb.WriteString("\n\n  --- Tako DX Demo: Native BubbleTea Component ---\n\n")
	sb.WriteString("  This component receives raw tea.Msg directly from the adapter,\n")
	sb.WriteString("  bypassing the router because it implements contracts.NativeComponent.\n\n")
	
	sb.WriteString(fmt.Sprintf("  Typed Text: %s\n\n", n.text))
	sb.WriteString("  (Type something, press ctrl+c to quit)\n")
	
	return sb.String()
}

func (n *NativeInput) RegisterKeys(keys contracts.KeyManager) {
	// We intentionally leave this blank or only register global keys.
	// Since we don't bind regular characters here, the Router will NOT consume them.
	// Thus, the adapter will forward them to UpdateNative().
}

// UpdateNative implements contracts.NativeComponent
func (n *NativeInput) UpdateNative(msg any) (any, any) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "backspace", "delete":
			if len(n.text) > 0 {
				n.text = n.text[:len(n.text)-1]
			}
		case "space":
			n.text += " "
		default:
			// Just a simple simulation: add printable chars
			if len(s) == 1 {
				n.text += s
			}
		}
	}
	
	return n, nil
}
