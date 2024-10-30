package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateCurrentTime(headTitle *widget.Label) {
	for {
		headTitle.SetText(time.Now().Format("Time: 15:04:05"))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(400, 200))
	headTitle := widget.NewLabel("Time: --:--:--")
	w.SetContent(headTitle)
	go func() {
		updateCurrentTime(headTitle)
	}()
	w.ShowAndRun()
}
