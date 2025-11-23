package main

import (
	"fyne.io/fyne/v2"

func main() {
	// App erstellen
	myApp := app.New()
	myWindow := myApp.NewWindow("Meine erste Fyne App")

	// Widget erstellen
	hello := widget.NewLabel("Hallo Fyne!")

	// Widget im Fenster anzeigen
	myWindow.SetContent(hello)

	// Fenster größe setzen und anzeigen
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
