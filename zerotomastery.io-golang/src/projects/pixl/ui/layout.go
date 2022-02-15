package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuidSwatches(app)

	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker)

	app.PixlWindow.SetContent(appLayout)
}