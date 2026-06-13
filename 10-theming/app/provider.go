package app

import (
	"github.com/gettako/tako/contracts"
	"github.com/gettako/tako/pkg/foundation"
)

type ThemeProvider struct {}

func (p *ThemeProvider) Register(app *foundation.Application) error { 
	// Setup languages
	app.Context().Lang().Register("en", map[string]string{"greeting": "Hello, World!"})
	app.Context().Lang().Register("id", map[string]string{"greeting": "Halo, Dunia!"})
	app.Context().Lang().SetLocale("en")
	
	// Setup themes
	app.Context().Theme().Register("dark", map[string]string{"bg": "0", "fg": "15", "accent": "42"})
	app.Context().Theme().Register("light", map[string]string{"bg": "15", "fg": "0", "accent": "21"})
	app.Context().Theme().Use("dark")
	
	return nil 
}

func (p *ThemeProvider) Boot(app *foundation.Application) error {
	ctx := app.Context()

	app.Stack().Push("base")
	var overlayMgr contracts.OverlayManager
	if err := app.Make(&overlayMgr); err == nil {
		overlayMgr.ShowComponent(&Dashboard{ctx: ctx})
	}
	return nil
}
