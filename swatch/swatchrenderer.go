package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (s *SwatchRenderer) MinSize() fyne.Size {
	return s.square.MinSize()
}

func (s *SwatchRenderer) Layout(size fyne.Size) {
	s.objects[0].Resize(size)
}

func (s *SwatchRenderer) Refresh() {
	s.Layout(fyne.NewSize(20, 20))
	s.square.FillColor = s.parent.Color

	if s.parent.Selected {
		s.square.StrokeWidth = 3
		s.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		s.objects[0] = &s.square
	} else {
		s.square.StrokeWidth = 0
		s.objects[0] = &s.square
	}

	canvas.Refresh(s.parent)
}

func (s *SwatchRenderer) Objects() []fyne.CanvasObject {
	return s.objects
}

func (s *SwatchRenderer) Destroy() {}
