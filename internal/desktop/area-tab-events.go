package desktop

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/desktop/uigen"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/constants"
	"github.com/martirosharutyunyan/unicode-converter/internal/modules/services"
)

func AreaTabEvents(ui *uigen.UIUnicodeConverterMainWindow) {
	formatConverterService := services.NewFormatConverterService()

	ui.UnicodeArea.ConnectTextChanged(OnChangeWrapper(ui, formatConverterService, constants.UNICODE))

	ui.AnsiArea.ConnectTextChanged(OnChangeWrapper(ui, formatConverterService, constants.ANSI))
}

func OnChangeWrapper(ui *uigen.UIUnicodeConverterMainWindow, formatConverterService services.IFormatConverterService, encoding constants.EncodingEnum) func() {
	switch encoding {
	case constants.UNICODE:
		return func() {
			UnicodeAreaOnChange(ui, formatConverterService)
		}
	default:
		return func() {
			AnsiAreaOnChange(ui, formatConverterService)
		}
	}
}

func UnicodeAreaOnChange(ui *uigen.UIUnicodeConverterMainWindow, formatConverterService services.IFormatConverterService) {
	output := formatConverterService.UnicodeToAnsi(ui.UnicodeArea.ToPlainText())
	ui.AnsiArea.DisconnectTextChanged()
	ui.AnsiArea.SetText(output)
	ui.AnsiArea.ConnectTextChanged(OnChangeWrapper(ui, formatConverterService, constants.ANSI))
}

func AnsiAreaOnChange(ui *uigen.UIUnicodeConverterMainWindow, formatConverterService services.IFormatConverterService) {
	output := formatConverterService.AnsiToUnicode(ui.AnsiArea.ToPlainText())
	ui.UnicodeArea.DisconnectTextChanged()
	ui.UnicodeArea.SetText(output)
	ui.UnicodeArea.ConnectTextChanged(OnChangeWrapper(ui, formatConverterService, constants.UNICODE))
}
