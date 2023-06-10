package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"local/layout"
)

var (
	data = [][]string{
		{"top left", "top right"},
		{"bottom left", "bottom right"},
	}
	content *fyne.Container
)

func main() {
	a := app.New()
	w := a.NewWindow("demo")
	w.Resize(fyne.NewSize(800, 500))
	table := Table()
	btn := widget.NewButton("Reset", func() {
		fmt.Println("reset clicked")
		data = [][]string{}
		table.Refresh()
	})
	btn2 := widget.NewButton("Recreate", func() {
		fmt.Println("recreate clicked")
		data = [][]string{}
		t := Table()
		content.Objects[2] = t
		content.Refresh()
	})
	c := container.New(&layout.LocalLayout{}, btn, btn2, table)
	content = c
	w.SetContent(c)
	w.ShowAndRun()
}

func Table() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*widget.Label)
			l.SetText(data[i.Row][i.Col])
		})
	return table
}
