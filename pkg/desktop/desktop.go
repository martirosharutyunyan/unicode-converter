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
	inputEntry := widget.NewEntry()

	openFileButton := widget.NewButton("Open File", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if uri != nil {
				defer uri.Close()
				inputEntry.SetText((uri.URI()).Path())
			}
		}, window)
	})

	// Create a simple UI layout
	content := container.NewVBox(
		openFileButton,
		inputEntry,
	)
	content.Resize(fyne.NewSize(100, 30))

	window.SetContent(content)
	window.Show()

	converterApp.Run()
}
