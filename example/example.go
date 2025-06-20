package main

import (
	"os"
	"strconv"

	"gioui.org/app"
	l "gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/nkrul/giox"
	xmat "github.com/nkrul/giox/material"
)

var (
	editor              widget.Editor
	checkbox            widget.Bool
	combo               giox.Combo
	comboSelectButton   widget.Clickable
	comboUnselectButton widget.Clickable
)

func main() {
	combo = giox.MakeCombo(
		[]string{
			"Option A",
			"Option B",
		},
		"select an option")

	run := func() {
		w := new(app.Window)
		loop(w)
	}

	go run()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			mainWindow(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}

func mainWindow(gtx l.Context, th *material.Theme) {

	for comboSelectButton.Clicked(gtx) {
		combo.SelectItem("Option B")
	}

	for comboUnselectButton.Clicked(gtx) {
		combo.Unselect()
	}

	children := []l.FlexChild{
		xmat.RigidSection(th, "giox Example"),
		xmat.RigidEditor(th, "Editor example", "<Insert some text here>", &editor),
		xmat.RigidCheckBox(th, "Checkbox example", &checkbox),
		l.Rigid(func(gtx l.Context) l.Dimensions {
			return xmat.Combo(th, &combo).Layout(gtx)
		}),
		xmat.RigidSeparator(th, &giox.Separator{}),
		xmat.RigidButton(th, "Force select Option B", &comboSelectButton),
		xmat.RigidButton(th, "Unselect", &comboUnselectButton),
		xmat.RigidLabel(th, strconv.FormatBool(checkbox.Value)),
	}

	if combo.HasSelected() {
		children = append(children, xmat.RigidLabel(th, combo.SelectedText()))
	}

	l.W.Layout(gtx, func(gtx l.Context) l.Dimensions {
		return l.Flex{Axis: l.Vertical}.Layout(gtx, children...)
	})
}
