package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/igor6629/pixel/apptype"
	"image"
	"image/color"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
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

func (p *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(p.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)

	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     p,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	p.renderer = renderer
	return renderer
}

func (p *PxCanvas) TryPan(previuosCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previuosCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		p.Pan(*previuosCoord, ev.PointEvent)
	}
}

func (p *PxCanvas) SetColor(c color.Color, x, y int) {
	if nrgba, ok := p.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}

	if rgba, ok := p.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}

	p.Refresh()
}

func (p *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := p.Bounds()

	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	pxSize := float32(p.PxSize)
	xOffset := p.CanvasOffset.X
	yOffset := p.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (p *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	p.PxCanvasConfig.PxCols = dimensions.Dx()
	p.PxCanvasConfig.PxRows = dimensions.Dy()

	p.PixelData = img
	p.reloadImage = true

	p.Refresh()
}

func (p *PxCanvas) NewDrawing(cols, rows int) {
	p.appState.SetFilePath("")
	p.PxCols = cols
	p.PxRows = rows

	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	p.LoadImage(pixelData)
}
