package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/igor6629/pixel/apptype"
	"github.com/igor6629/pixel/swatch"
	"github.com/igor6629/pixel/ui"
	"image/color"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("Pixel")

	state := apptype.State{
		BrushColor:     color.NRGBA{0, 0, 0, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	pixelApp.Settings().SetTheme(theme.DarkTheme())
	appInit.PixelWindow.ShowAndRun()
}
