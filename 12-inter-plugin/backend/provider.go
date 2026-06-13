package backend

import (
	"context"
	"fmt"
	"time"

	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

// BackendProvider simulates an independent plugin that provides data via RPC
// and reacts to events, but has no UI components of its own.
type BackendProvider struct{}

func (p *BackendProvider) Register(_ *foundation.Application) error { return nil }

func (p *BackendProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	// 1. Listen for Events from other plugins
	ctx.Events().On("frontend:ping", func(e contracts.Event) {
		msg := e.Data.(string)
		
		// Respond back
		response := fmt.Sprintf("Backend received: '%s'", msg)
		ctx.Events().Publish("backend:pong", response)
	})

	// 2. Expose an RPC endpoint for synchronous calls
	ctx.RPC().Route("weather:get").Handle(func(_ context.Context, req contracts.RPCRequest) (contracts.RPCResponse, error) {
		time.Sleep(500 * time.Millisecond) // Simulate network delay
		
		location, ok := req.Payload.(string)
		if !ok || location == "" {
			return contracts.RPCResponse{}, fmt.Errorf("location must be a string")
		}

		// Return mock weather data
		weather := fmt.Sprintf("It's sunny and 28°C in %s", location)
		return contracts.RPCResponse{Data: weather}, nil
	})

	return nil
}
