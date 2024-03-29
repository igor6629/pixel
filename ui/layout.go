package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetupMenu(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixelCanvas)

	app.PixelWindow.SetContent(appLayout)
}
