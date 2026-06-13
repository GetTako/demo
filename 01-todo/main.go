package main

import (
	"os"
	"log"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/01-todo/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()

	// Register providers explicitly
	if err := takoApp.RegisterProviders(&app.TodoProvider{}); err != nil {
		log.Fatalf("Failed to register providers: %v", err)
	}

	// Wire up the Bubble Tea rendering adapter with our Layout
	takoApp.WithRenderer(bubbletea.NewAdapter(takoApp.Context(), &app.Layout{}))

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
