package main

import (
	"log"
	"os"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/02-dashboard/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()

	if err := takoApp.RegisterProviders(&app.DashboardProvider{}); err != nil {
		log.Fatalf("Failed to register providers: %v", err)
	}

	takoApp.Mount(bubbletea.NewAdapter(takoApp.Context(), nil))

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
