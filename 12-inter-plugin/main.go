package main

import (
	"log"
	"os"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/12-inter-plugin/backend"
	"github.com/gettako/tako/demo/12-inter-plugin/frontend"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	app := tako.NewApp()

	// Register multiple independent providers (simulating distinct plugins)
	app.RegisterProviders(
		&backend.BackendProvider{},
		&frontend.FrontendProvider{},
	)

	// Boot the TUI layout from frontend
	adapter := bubbletea.NewAdapter(app.Context(), nil)
	app.Mount(adapter)

	if err := tako.Run(app, os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
