package desktop

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/desktop/uigen"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/services"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/utils"
	"github.com/therecipe/qt/widgets"
	"os"
)

func FileTabEvents(ui *uigen.UIUnicodeConverterMainWindow) {
	convertFileService := services.NewConvertFileService()
	convertDirService := services.NewConvertDirService()

	var isFile bool

	ui.OutputSource.SetReadOnly(true)

	ui.AnsiToUnicode.SetChecked(true)
	ui.OpenFileButton.ConnectClicked(func(_ bool) {
		options := widgets.QFileDialog__ReadOnly
		file := widgets.QFileDialog_GetOpenFileName(nil, "Open File", "", "", "", options)
		if file != "" {
			ui.InputSource.SetText(file)
			outputFile, err := utils.GenCopyPath(file)
			if err != nil {
				AlertError(err)
			}
			ui.OutputSource.SetText(outputFile)
		}
		isFile = true
	})

	ui.OpenFolderButton.ConnectClicked(func(checked bool) {
		options := widgets.QFileDialog__ReadOnly
		dirName := widgets.QFileDialog_GetExistingDirectory(nil, "Open Dir", "", options)
		if dirName != "" {
			ui.InputSource.SetText(dirName)
			outputFile, err := utils.GenCopyPath(dirName)
			if err != nil {
				AlertError(err)
			}
			ui.OutputSource.SetText(outputFile)
		}
		isFile = false
	})

	ui.SaveButton.ConnectClicked(func(checked bool) {
		inputSource := ui.InputSource.Text()
		outputSource := ui.OutputSource.Text()

		if isFile {
			err := convertFileService.ConvertFile(inputSource, outputSource, ui.UnicodeToAnsi.IsChecked())
			if err != nil {
				AlertError(err)
			}
		} else {
			err := os.Mkdir(outputSource, 0777)
			if err != nil {
				AlertError(err)
			}
			err = convertDirService.ConvertDir(inputSource, outputSource, ui.UnicodeToAnsi.IsChecked())
			if err != nil {
				AlertError(err)
			}
		}
	})
}
