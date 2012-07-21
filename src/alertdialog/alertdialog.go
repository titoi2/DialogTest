package alertdialog

import (
	"github.com/visualfc/go-ui/ui"
)

type AlertDialog struct {
	*ui.Dialog
	lbl *ui.Label
	btn *ui.Button
}

func NewAlertDialog() *AlertDialog {
	return new(AlertDialog).Init()
}

func (p *AlertDialog) Init() *AlertDialog {
	p.Dialog = ui.NewDialog()
	p.SetWindowTitle("Alert Dialog")

	vbox := ui.NewVBoxLayout()

	p.SetLayout(vbox)

	btn := ui.NewButton()
	btn.SetText("OK")
	btn.OnClicked(func() {
		p.Close()
	})

	p.btn = btn

	lbl := ui.NewLabel()
	p.lbl = lbl

	vbox.AddWidget(lbl)
//	vbox.AddStretch(0)
	vbox.AddWidget(btn)

	return p
}

func (p *AlertDialog) SetTitle(s string) {
	p.SetWindowTitle("Alert Dialog")
}
func (p *AlertDialog) SetMessage(s string) {
	p.lbl.SetText(s)
}
