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
	label := widget.NewLabel("Input file path")

	content := container.NewVBox(
		openFileButton,
		label,
		inputEntry,
	)
	content.Resize(fyne.NewSize(100, 30))

	tabs := container.NewAppTabs(
		container.NewTabItem("Convert file or directory", content),
		container.NewTabItem("Convert Text", widget.NewLabel("Convert text")))

	tabs.SetTabLocation(container.TabLocationLeading)

	window.SetContent(tabs)
	window.Show()

	converterApp.Run()
}
