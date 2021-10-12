package main

import (
	_ "embed"

	"github.com/NWBY/beluga/lib"
	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "beluga",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(lib.NewBeluga())
	app.Run()
}
