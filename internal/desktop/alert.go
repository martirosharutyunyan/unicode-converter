package desktop

import "github.com/therecipe/qt/widgets"

func AlertError(err error) {
	msgBox := widgets.NewQMessageBox(nil)
	msgBox.SetIcon(widgets.QMessageBox__Critical)
	msgBox.SetText("Error")
	msgBox.SetInformativeText(err.Error())
	msgBox.SetStandardButtons(widgets.QMessageBox__Ok)
	msgBox.Exec()
}

func AlertOk(message string) {
	msgBox := widgets.NewQMessageBox(nil)
	msgBox.SetIcon(widgets.QMessageBox__Information)
	msgBox.SetText("OK")
	msgBox.SetInformativeText(message)
	msgBox.SetStandardButtons(widgets.QMessageBox__Ok)
	msgBox.Exec()
}
