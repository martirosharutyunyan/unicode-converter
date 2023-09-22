package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/martirosharutyunyan/unicode-converter/pkg/internal/modules/services"
	static_errors "github.com/martirosharutyunyan/unicode-converter/pkg/internal/modules/static-errors"
	"github.com/martirosharutyunyan/unicode-converter/pkg/internal/modules/utils"
)

func ConvertSourceContainer(window fyne.Window) *fyne.Container {
	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("Output Source")
	inputSource := ""
	isUnicodeToAnsi := false

	openFileButton := widget.NewButton("Open File", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if uri != nil {
				defer uri.Close()
				inputSource = (uri.URI()).Path()
				outputSource, err := utils.GenCopyPath(inputSource)
				if err != nil {
					dialog.NewError(err, window)
				}
				outputEntry.SetText(outputSource)
			}
		}, window)
	})
	openDirButton := widget.NewButton("Open Dir", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				inputSource = uri.Path()
				outputSource, err := utils.GenCopyPath(inputSource)
				if err != nil {
					dialog.NewError(err, window).Show()
				}
				outputEntry.SetText(outputSource)
			}
		}, window)
	})
	buttonsContainer := container.NewHBox(openFileButton, openDirButton)

	radioGroup := widget.NewRadioGroup([]string{"Ansi to Unicode", "Unicode to Ansi"}, func(s string) {
		if s == "Ansi to Unicode" {
			isUnicodeToAnsi = false
		}
	})

	saveButton := widget.NewButton("Save", func() {
		if inputSource == "" {
			dialog.NewError(static_errors.TRANSFER_SOURCE_ERR, window).Show()
			return
		}
		convertFileService := services.NewConvertFileService()
		err := convertFileService.ConvertFile(inputSource, outputEntry.Text, isUnicodeToAnsi)
		if err != nil {
			dialog.NewError(err, window).Show()
		}
	})

	convertFileContainer := container.NewVBox(
		buttonsContainer,
		radioGroup,
		outputEntry,
		layout.NewSpacer(),
		saveButton,
	)

	return convertFileContainer
}
