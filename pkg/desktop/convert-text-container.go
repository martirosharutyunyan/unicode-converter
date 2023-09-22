package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/martirosharutyunyan/unicode-converter/pkg/internal/modules/services"
)

func ConvertTextContainer(window fyne.Window) *fyne.Container {
	formatConverterService := services.NewFormatConverterService()

	unicodeLabel := widget.NewLabel("Unicode")
	unicodeEntry := widget.NewEntry()

	ansiLabel := widget.NewLabel("Ansi")
	ansiEntry := widget.NewEntry()

	unicodeEntry.OnChanged = func(s string) {
		output := formatConverterService.UnicodeToAnsi(s)
		ansiEntry.SetText(output)
		window.Clipboard().SetContent(output)
	}

	ansiEntry.OnChanged = func(s string) {
		output := formatConverterService.AnsiToUnicode(s)
		unicodeEntry.SetText(output)
		window.Clipboard().SetContent(output)
	}

	unicodeContainer := container.NewWithoutLayout(unicodeLabel, unicodeEntry)
	ansiContainer := container.NewWithoutLayout(ansiLabel, ansiEntry)

	unicodeEntry.Resize(fyne.NewSize(470, 500))
	ansiEntry.Resize(fyne.NewSize(470, 500))

	unicodeLabel.Move(fyne.NewPos(190, 0))
	ansiLabel.Move(fyne.NewPos(640, 0))

	unicodeEntry.Move(fyne.NewPos(0, 40))
	ansiEntry.Move(fyne.NewPos(420, 40))

	textConverterContainer := container.NewHBox(
		unicodeContainer,
		ansiContainer,
	)

	return textConverterContainer
}
