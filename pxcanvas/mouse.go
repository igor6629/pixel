package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/igor6629/pixel/pxcanvas/brush"
)

func (p *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	p.scale(int(ev.Scrolled.DY))
	p.Refresh()
}

func (p *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := p.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(p.appState, p, ev)
		cursor := brush.Cursor(p.PxCanvasConfig, p.appState.BrushType, ev, *x, *y)
		p.renderer.SetCursor(cursor)
	} else {
		p.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}

	p.TryPan(p.mouseState.previousCoord, ev)
	p.Refresh()
	p.mouseState.previousCoord = &ev.PointEvent
}

func (p *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (p *PxCanvas) MouseOut()                      {}

func (p *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(p.appState, p, ev)
}

func (p *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}
