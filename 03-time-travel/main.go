package main

import (
	"os"
	"log"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/03-time-travel/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()

	if err := takoApp.RegisterProviders(&app.TimeTravelProvider{}); err != nil {
		log.Fatalf("Failed to register providers: %v", err)
	}

	takoApp.WithRenderer(bubbletea.NewAdapter(takoApp.Context(), &app.Layout{}))

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
