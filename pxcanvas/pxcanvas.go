package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/igor6629/pixel/apptype"
	"image"
	"image/color"
)

type PxCanvasMouseState struct {
	previousCord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (p *PxCanvas) Bounds() image.Rectangle {
	x0 := int(p.CanvasOffset.X)
	y0 := int(p.CanvasOffset.Y)

	x1 := int(p.PxCols*p.PxSize + int(p.CanvasOffset.X))
	y1 := int(p.PxRows*p.PxSize + int(p.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) && pos.X < float32(bounds.Max.X) && pos.Y >= float32(bounds.Min.Y) && pos.Y < float32(bounds.Max.Y) {
		return true
	}

	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			img.Set(j, i, c)
		}
	}

	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}

	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)

	return pxCanvas
}
