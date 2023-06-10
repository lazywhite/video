package layout

import (
	"fyne.io/fyne/v2"
)

type LocalLayout struct {
}

func (l LocalLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	objects[0].Resize(fyne.NewSize(100, 50))
	objects[0].Move(fyne.NewPos(0, 0))
	objects[1].Resize(fyne.NewSize(100, 50))
	objects[1].Move(fyne.NewPos(100, 0))
	objects[2].Resize(fyne.NewSize(800, 300))
	objects[2].Move(fyne.NewPos(0, 50))
}

func (l LocalLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 400)
}
