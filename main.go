package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

var names = []string{
	"John", "Doe",
	"Jane", "Smith",
}

func main() {
	log.Print("Hello World!")
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(400, 200))

	headTitle := widget.NewLabel("Hello World!")
	w.SetContent(headTitle)
	go func() {
		for _, name := range names {
			headTitle.SetText(name)
			time.Sleep(1 * time.Second)
		}
	}()
	w.ShowAndRun()
}
