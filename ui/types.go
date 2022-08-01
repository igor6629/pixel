package ui

import (
	"fyne.io/fyne/v2"
	"github.com/igor6629/pixel/apptype"
	"github.com/igor6629/pixel/pxcanvas"
	"github.com/igor6629/pixel/swatch"
)

type AppInit struct {
	PixelCanvas *pxcanvas.PxCanvas
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
