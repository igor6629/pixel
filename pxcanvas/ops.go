package pxcanvas

import "fyne.io/fyne/v2"

func (p *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	p.CanvasOffset.X += xDiff
	p.CanvasOffset.Y += yDiff
	p.Refresh()
}

func (p *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		p.PxSize++
	case direction < 0:
		if p.PxSize > 2 {
			p.PxSize--
		}
	default:
		p.PxSize = 10
	}
}
