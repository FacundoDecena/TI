package main

import (
	"TI.exe/HammingCodification"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "TI",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(preHamming)
	app.Bind(HammingCodification.Hamming7)
	app.Bind(basic)
	_ = app.Run()
}
