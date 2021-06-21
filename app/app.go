package main

import (
	"DangoTranslator/app/ui/settin"
	"github.com/therecipe/qt/widgets"
	"os"
)

func init() {
}

func main() {

	App := widgets.NewQApplication(len(os.Args), os.Args)

	SettingUI := settin.UI()
	SettingUI.Show()

	App.Exec()
}
