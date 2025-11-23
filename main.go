package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()                  // neue Fyne-Anwendung erstellen
	w := a.NewWindow("Hello World") // neues Fenster mit Titel

	w.SetContent(widget.NewLabel("Hello World!")) // Label als Inhalt setzen
	fmt.Println("Running")

	w.ShowAndRun() // Fenster anzeigen und Anwendung starten
}
