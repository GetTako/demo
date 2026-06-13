package main

import (
	"log"
	"os"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/11-state/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()
	takoApp.RegisterProviders(&app.StateProvider{})
	adapter := bubbletea.NewAdapter(takoApp.Context(), &app.MainLayout{})
	takoApp.WithRenderer(adapter)

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
