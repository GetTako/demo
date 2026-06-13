package main

import (
	"os"
	"log"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/04-mouse/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()

	if err := takoApp.RegisterProviders(&app.MouseProvider{}); err != nil {
		log.Fatalf("Failed to register providers: %v", err)
	}

	takoApp.WithRenderer(bubbletea.NewAdapter(takoApp.Context(), &app.Layout{}))

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
