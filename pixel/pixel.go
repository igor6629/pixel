package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/igor6629/pixel/apptype"
	"github.com/igor6629/pixel/pxcanvas"
	"github.com/igor6629/pixel/swatch"
	"github.com/igor6629/pixel/ui"
	"image/color"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("Pixel")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixelCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       20,
	}

	pixelCanvas := pxcanvas.NewPxCanvas(&state, pixelCanvasConfig)

	appInit := ui.AppInit{
		PixelCanvas: pixelCanvas,
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	pixelApp.Settings().SetTheme(theme.DarkTheme())
	appInit.PixelWindow.CenterOnScreen()
	appInit.PixelWindow.ShowAndRun()
}
