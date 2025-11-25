package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	items := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9", "Item 10"}
	len_items := len(items)
	fmt.Print(" number of items: ", len_items)

	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(items[i])
		})

	scroll := container.NewVScroll(list)
	w.SetContent(scroll)

	fmt.Println("Running")
	w.Resize(fyne.NewSize(400, 200))

	w.ShowAndRun()
}
