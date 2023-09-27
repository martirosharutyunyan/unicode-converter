// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
)

type UIUnicodeConverterMainWindow struct {
	Centralwidget *widgets.QWidget
	TextConverter *widgets.QTabWidget
	TextConverter2 *widgets.QWidget
	TextEdit *widgets.QTextEdit
	TextEdit2 *widgets.QTextEdit
	Label *widgets.QLabel
	Label2 *widgets.QLabel
	SourceConverter *widgets.QWidget
	RadioButton3 *widgets.QRadioButton
	RadioButton4 *widgets.QRadioButton
	LineEdit *widgets.QLineEdit
	PushButton *widgets.QPushButton
	PushButton2 *widgets.QPushButton
	PushButton3 *widgets.QPushButton
}

func (this *UIUnicodeConverterMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 797, 487))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.TextConverter = widgets.NewQTabWidget(this.Centralwidget)
	this.TextConverter.SetObjectName("TextConverter")
	this.TextConverter.SetGeometry(core.NewQRect4(0, 0, 821, 621))
	this.TextConverter.SetMinimumSize(core.NewQSize2(211, 51))
	this.TextConverter2 = widgets.NewQWidget(this.TextConverter, core.Qt__Widget)
	this.TextConverter2.SetObjectName("TextConverter2")
	this.TextEdit = widgets.NewQTextEdit(this.TextConverter2)
	this.TextEdit.SetObjectName("TextEdit")
	this.TextEdit.SetGeometry(core.NewQRect4(10, 60, 361, 371))
	this.TextEdit2 = widgets.NewQTextEdit(this.TextConverter2)
	this.TextEdit2.SetObjectName("TextEdit2")
	this.TextEdit2.SetGeometry(core.NewQRect4(400, 60, 361, 371))
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
	this.RadioButton3 = widgets.NewQRadioButton(this.SourceConverter)
	this.RadioButton3.SetObjectName("RadioButton3")
	this.RadioButton3.SetGeometry(core.NewQRect4(20, 60, 131, 22))
	this.RadioButton4 = widgets.NewQRadioButton(this.SourceConverter)
	this.RadioButton4.SetObjectName("RadioButton4")
	this.RadioButton4.SetGeometry(core.NewQRect4(20, 90, 131, 22))
	this.LineEdit = widgets.NewQLineEdit(this.SourceConverter)
	this.LineEdit.SetObjectName("LineEdit")
	this.LineEdit.SetGeometry(core.NewQRect4(20, 120, 441, 32))
	this.PushButton = widgets.NewQPushButton(this.SourceConverter)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetGeometry(core.NewQRect4(690, 410, 88, 34))
	this.PushButton2 = widgets.NewQPushButton(this.SourceConverter)
	this.PushButton2.SetObjectName("PushButton2")
	this.PushButton2.SetGeometry(core.NewQRect4(20, 10, 88, 34))
	this.PushButton3 = widgets.NewQPushButton(this.SourceConverter)
	this.PushButton3.SetObjectName("PushButton3")
	this.PushButton3.SetGeometry(core.NewQRect4(120, 10, 88, 34))
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
	this.RadioButton3.SetText(_translate("MainWindow", "Unicode to Ansi", "", -1))
	this.RadioButton4.SetText(_translate("MainWindow", "Ansi to Unicode", "", -1))
	this.LineEdit.SetPlaceholderText(_translate("MainWindow", "Output source", "", -1))
	this.PushButton.SetText(_translate("MainWindow", "Save", "", -1))
	this.PushButton2.SetText(_translate("MainWindow", "Open file", "", -1))
	this.PushButton3.SetText(_translate("MainWindow", "Open Folder", "", -1))
	this.TextConverter.SetTabText(this.TextConverter.IndexOf(this.SourceConverter), _translate("MainWindow", "Convert Source", "", -1))
}
