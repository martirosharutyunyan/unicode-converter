package desktop

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/desktop/uigen"
	"github.com/therecipe/qt/widgets"
	"os"
)

func Run() {

	ui := new(uigen.UIUnicodeConverterMainWindow)
	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	qApp.SetApplicationName("Converter")
	mainWindow := widgets.NewQMainWindow(nil, 0)
	ui.SetupUI(mainWindow)
	mainWindow.SetWindowTitle("Converter")

	AreaTabEvents(ui)
	FileTabEvents(ui)

	mainWindow.Show()
	qApp.Exec()
}
