package main

import (
	"log"
	"os"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/10-theming/app"
	"github.com/gettako/tako/pkg/adapter/bubbletea"
)

func main() {
	takoApp := tako.NewApp()
	takoApp.RegisterProviders(&app.ThemeProvider{})
	adapter := bubbletea.NewAdapter(takoApp.Context(), nil)
	takoApp.Mount(adapter)

	if err := tako.Run(takoApp, os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
