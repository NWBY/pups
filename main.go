package main

import (
	_ "embed"

	"github.com/NWBY/pups/lib"
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
		Title:  "pups",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(lib.NewPups())
	app.Run()
}
