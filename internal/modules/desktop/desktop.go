package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func App() {
	converterApp := app.New()
	window := converterApp.NewWindow("Converter")

	window.Resize(fyne.NewSize(1000, 500))

	filePathLabel := widget.NewLabel("Selected File: ")

	// Create a button to open the file dialog
	openFileButton := widget.NewButton("Open File", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			defer uri.Close()

		}, window)
	})

	// Create a simple UI layout
	content := container.NewVBox(
		openFileButton,
		filePathLabel,
	)
	window.SetContent(content)
	window.Show()

	converterApp.Run()
}
