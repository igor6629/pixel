package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (p *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	p.scale(int(ev.Scrolled.DY))
	p.Refresh()
}

func (p *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	p.TryPan(p.mouseState.previousCoord, ev)
	p.Refresh()
	p.mouseState.previousCoord = &ev.PointEvent
}

func (p *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (p *PxCanvas) MouseOut()                      {}
