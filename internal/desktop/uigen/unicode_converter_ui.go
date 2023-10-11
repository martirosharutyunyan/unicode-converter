// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
)

type UIUnicodeConverterMainWindow struct {
	Centralwidget *widgets.QWidget
	TextConverter *widgets.QTabWidget
	TextConverter2 *widgets.QWidget
	UnicodeArea *widgets.QTextEdit
	AnsiArea *widgets.QTextEdit
	Label *widgets.QLabel
	Label2 *widgets.QLabel
	SourceConverter *widgets.QWidget
	UnicodeToAnsi *widgets.QRadioButton
	AnsiToUnicode *widgets.QRadioButton
	InputSource *widgets.QLineEdit
	SaveButton *widgets.QPushButton
	OpenFileButton *widgets.QPushButton
	OpenFolderButton *widgets.QPushButton
	OutputSource *widgets.QLineEdit
}

func (this *UIUnicodeConverterMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 797, 487))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.TextConverter = widgets.NewQTabWidget(this.Centralwidget)
	this.TextConverter.SetObjectName("TextConverter")
	this.TextConverter.SetGeometry(core.NewQRect4(-10, 0, 821, 621))
	this.TextConverter.SetMinimumSize(core.NewQSize2(211, 51))
	this.TextConverter2 = widgets.NewQWidget(this.TextConverter, core.Qt__Widget)
	this.TextConverter2.SetObjectName("TextConverter2")
	this.UnicodeArea = widgets.NewQTextEdit(this.TextConverter2)
	this.UnicodeArea.SetObjectName("UnicodeArea")
	this.UnicodeArea.SetGeometry(core.NewQRect4(10, 60, 361, 371))
	this.AnsiArea = widgets.NewQTextEdit(this.TextConverter2)
	this.AnsiArea.SetObjectName("AnsiArea")
	this.AnsiArea.SetGeometry(core.NewQRect4(400, 60, 361, 371))
	this.Label = widgets.NewQLabel(this.TextConverter2, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.Label.SetGeometry(core.NewQRect4(150, 20, 71, 18))
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetPointSize(12)
	this.Label.SetFont(font)
	this.Label2 = widgets.NewQLabel(this.TextConverter2, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.Label2.SetGeometry(core.NewQRect4(550, 20, 61, 18))
	font = gui.NewQFont()
	font.SetPointSize(12)
	this.Label2.SetFont(font)
	this.TextConverter.AddTab(this.TextConverter2, "")
	this.SourceConverter = widgets.NewQWidget(this.TextConverter, core.Qt__Widget)
	this.SourceConverter.SetObjectName("SourceConverter")
	this.UnicodeToAnsi = widgets.NewQRadioButton(this.SourceConverter)
	this.UnicodeToAnsi.SetObjectName("UnicodeToAnsi")
	this.UnicodeToAnsi.SetGeometry(core.NewQRect4(20, 60, 131, 22))
	this.AnsiToUnicode = widgets.NewQRadioButton(this.SourceConverter)
	this.AnsiToUnicode.SetObjectName("AnsiToUnicode")
	this.AnsiToUnicode.SetGeometry(core.NewQRect4(20, 90, 131, 22))
	this.InputSource = widgets.NewQLineEdit(this.SourceConverter)
	this.InputSource.SetObjectName("InputSource")
	this.InputSource.SetGeometry(core.NewQRect4(20, 120, 721, 32))
	this.SaveButton = widgets.NewQPushButton(this.SourceConverter)
	this.SaveButton.SetObjectName("SaveButton")
	this.SaveButton.SetGeometry(core.NewQRect4(690, 410, 88, 34))
	this.OpenFileButton = widgets.NewQPushButton(this.SourceConverter)
	this.OpenFileButton.SetObjectName("OpenFileButton")
	this.OpenFileButton.SetGeometry(core.NewQRect4(20, 10, 88, 34))
	this.OpenFolderButton = widgets.NewQPushButton(this.SourceConverter)
	this.OpenFolderButton.SetObjectName("OpenFolderButton")
	this.OpenFolderButton.SetGeometry(core.NewQRect4(120, 10, 88, 34))
	this.OutputSource = widgets.NewQLineEdit(this.SourceConverter)
	this.OutputSource.SetObjectName("OutputSource")
	this.OutputSource.SetGeometry(core.NewQRect4(20, 160, 721, 32))
	this.TextConverter.AddTab(this.SourceConverter, "")
	MainWindow.SetCentralWidget(this.Centralwidget)


    this.RetranslateUi(MainWindow)
	this.TextConverter.SetCurrentIndex(1)
}

func (this *UIUnicodeConverterMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.Label.SetText(_translate("MainWindow", "Unicode", "", -1))
	this.Label2.SetText(_translate("MainWindow", "Ansi", "", -1))
	this.TextConverter.SetTabText(this.TextConverter.IndexOf(this.TextConverter2), _translate("MainWindow", "Convert Text", "", -1))
	this.UnicodeToAnsi.SetText(_translate("MainWindow", "Unicode to Ansi", "", -1))
	this.AnsiToUnicode.SetText(_translate("MainWindow", "Ansi to Unicode", "", -1))
	this.InputSource.SetPlaceholderText(_translate("MainWindow", "Input source", "", -1))
	this.SaveButton.SetText(_translate("MainWindow", "Save", "", -1))
	this.OpenFileButton.SetText(_translate("MainWindow", "Open file", "", -1))
	this.OpenFolderButton.SetText(_translate("MainWindow", "Open Folder", "", -1))
	this.OutputSource.SetPlaceholderText(_translate("MainWindow", "Output source", "", -1))
	this.TextConverter.SetTabText(this.TextConverter.IndexOf(this.SourceConverter), _translate("MainWindow", "Convert Source", "", -1))
}
