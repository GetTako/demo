package main

import (
	"log"
	"os"

	"github.com/gettako/tako"
	"github.com/gettako/tako/demo/09-cli/commands"
)

func main() {
	app := tako.NewApp()

	// Register CLI commands. 
	// The app doesn't have a UI adapter (No BubbleTea)!
	// It will exclusively run as a CLI tool.
	app.Commands(&commands.HelloCommand{})

	// Pass arguments directly to Tako's CLI kernel
	if err := tako.Run(app, os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
