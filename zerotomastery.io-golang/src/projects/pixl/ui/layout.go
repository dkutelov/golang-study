package ui

func Setup(app *AppInit) {
	swatchesContainer := BuidSwatches(app)

	app.PixlWindow.SetContent(swatchesContainer)
}