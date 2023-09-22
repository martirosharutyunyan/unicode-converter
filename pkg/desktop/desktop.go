package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"os"
)

func App() {
	converterApp := app.New()
	window := converterApp.NewWindow("Converter")
	window.Resize(fyne.NewSize(1100, 600))

	sourceConverterTab := container.NewTabItem("Convert Source", ConvertSourceContainer(window))
	textConverterTab := container.NewTabItem("Convert Text", ConvertTextContainer(window))
	tabs := container.NewAppTabs(
		textConverterTab,
		sourceConverterTab,
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.Select(textConverterTab)

	window.SetContent(tabs)
	window.Show()

	converterApp.Run()
	os.Exit(0)
}
