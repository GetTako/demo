package main

import (
	"log"
	"os"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/07-communication/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()

	// Register our main feature provider
	takoApp.RegisterProviders(&app.CommunicationProvider{})

	// Initialize the Bubble Tea UI adapter
	adapter := bubbletea.NewAdapter(takoApp.Context(), nil)
	takoApp.Mount(adapter)

	// Run the application
	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
