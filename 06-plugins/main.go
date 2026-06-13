package main

import (
	"os"
	"log"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/06-plugins/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()

	// External plugin: has Manifest(), so it appears in `plugin:list`
	if err := takoApp.RegisterProviders(&app.DemoPlugin{}); err != nil {
		log.Fatalf("Failed to register providers: %v", err)
	}

	takoApp.WithRenderer(bubbletea.NewAdapter(takoApp.Context(), &app.Layout{}))

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
