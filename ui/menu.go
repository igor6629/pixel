package ui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/igor6629/pixel/util"
	"image"
	"image/png"
	"os"
	"strconv"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PixelCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixelWindow)
				return
			}

			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PixelWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save as...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PixelWindow)
				}
			}
			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PixelWindow)
				return
			}
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}

			if width <= 0 {
				return errors.New("must be > 0")
			}

			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0

				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixelWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixelWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}

				app.PixelCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixelWindow)
	})
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, e error) {
			if uri == nil {
				return
			} else {
				img, _, err := image.Decode(uri)
				if err != nil {
					dialog.ShowError(err, app.PixelWindow)
					return
				}

				app.PixelCanvas.LoadImage(img)
				app.State.SetFilePath(uri.URI().Path())
				imgColors := util.GetImageColors(img)
				i := 0

				for c := range imgColors {
					if i == len(app.Swatches) {
						break
					}

					app.Swatches[i].SetColor(c)
					i++
				}
			}
		}, app.PixelWindow)
	})
}

func BuildMenu(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File", BuildNewMenu(app), BuildOpenMenu(app), BuildSaveMenu(app), BuildSaveAsMenu(app))
}

func SetupMenu(app *AppInit) {
	menu := BuildMenu(app)
	mainMenu := fyne.NewMainMenu(menu)
	app.PixelWindow.SetMainMenu(mainMenu)
}
