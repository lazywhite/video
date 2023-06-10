package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	data = [][]string{
		{"top left", "top right"},
		{"bottom left", "bottom right"},
	}
)

func main() {
	a := app.New()
	w := a.NewWindow("demo")
	w.Resize(fyne.NewSize(800, 500))
	table := RightTable()
	//table := WrongTable()

	c := container.NewMax(table)
	w.SetContent(c)
	w.ShowAndRun()
}

func WrongTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 3
		},
		func() fyne.CanvasObject {
			return container.NewHBox(layout.NewSpacer())
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)
			switch i.Col {
			case 2:
				btn := widget.NewButton("Edit", func() {
					fmt.Printf("row: %d, col: %d\n", i.Row, i.Col)
				})
				btn.Importance = widget.HighImportance
				btn.Refresh()
				c.Objects[0] = btn
			default:
				l := widget.NewLabel("template")
				l.SetText(data[i.Row][i.Col])
				c.Objects[0] = l
				c.Refresh()
			}
		})
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 200)
	table.SetColumnWidth(1, 200)
	return table
}

func RightTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 3
		},
		func() fyne.CanvasObject {
			l := widget.NewLabel("template")
			b := widget.NewButton("Delete", func() {})
			b.Importance = widget.HighImportance
			return container.NewHBox(l, b)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)
			l := c.Objects[0].(*widget.Label)
			btn := c.Objects[1].(*widget.Button)
			switch i.Col {
			case 2:
				btn.OnTapped = func() {
					fmt.Printf("row: %d, col: %d\n", i.Row, i.Col)
				}
				btn.Show()
				l.Hide()
			default:
				l.SetText(data[i.Row][i.Col])
				l.Show()
				btn.Hide()
			}
		})
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 200)
	table.SetColumnWidth(1, 200)
	return table
}
