package main

import (
	"log"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/13-native-bubbles/app"
)

func main() {
	application := tako.NewApp()

	application.RegisterProviders(
		&app.Provider{},
	)

	if err := tako.Run(application); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
