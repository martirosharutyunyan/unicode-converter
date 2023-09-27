package desktop

import (
	"github.com/martirosharutyunyan/unicode-converter/internal/desktop/uigen"
	"github.com/therecipe/qt/widgets"
	"os"
)

func Run() {
	ui := new(uigen.UIUnicodeConverterMainWindow)
	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	mainWindow := widgets.NewQMainWindow(nil, 0)
	ui.SetupUI(mainWindow)

	mainWindow.Show()
	qApp.Exec()
}
