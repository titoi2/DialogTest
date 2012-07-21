package main

import (
	"./alertdialog"
	"github.com/visualfc/go-ui/ui"
)

var exit = make(chan bool)

func main() {
	ui.Main(func() {
		go ui_main()
		ui.Run()
		exit <- true
	})
}

func ui_main() {

	w := ui.NewWidget()
	w.SetWindowTitle(ui.Version())
	w.SetSizev(300, 200)
	defer w.Close()

	d := alertdialog.NewAlertDialog()
	d.SetMessage("Hello\nWelcome Go lang world!!")

	d.Open()
	defer d.Close()

	w.Show()
	<-exit
}
