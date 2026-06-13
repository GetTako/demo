// Package main is the entry point for the demo.
package main

import (
	"log"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/13-native-bubbles/app"
)

func main() {
	application := tako.NewApp()

	if err := application.RegisterProviders(
		&app.Provider{},
	); err != nil {
		log.Fatalf("Error registering providers: %v", err)
	}

	if err := tako.Run(application); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
