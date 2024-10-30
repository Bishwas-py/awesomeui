package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"log"
	"reflect"
)

type human struct {
	Name           string
	Age            int
	Profession     string
	Gender         string
	ArrivalDate    string
	DepartureDate  string
	Reason         string
	IsInterstellar bool
	Inhabitant     string
}

func main() {
	phobe := human{
		Name:           "Phobe",
		Age:            25,
		Profession:     "Software Engineer",
		Gender:         "Female",
		ArrivalDate:    "2021-09-01",
		DepartureDate:  "2021-09-30",
		Reason:         "Work",
		IsInterstellar: false,
		Inhabitant:     "Earth",
	}
	myApp := app.New()
	myWindow := myApp.NewWindow("Human Information")
	myWindow.Resize(fyne.NewSize(100, 400))
	humanDisplayUI := getHumanDisplayUI(phobe)
	content := container.New(layout.NewHBoxLayout())
	for _, pair := range humanDisplayUI {
		content.Add(container.New(layout.NewVBoxLayout(), pair.Label, pair.Text))
	}
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

type humanDisplayUI struct {
	Label *canvas.Text
	Text  *canvas.Text
}

func getHumanDisplayUI(human human) []humanDisplayUI {
	var pairs []humanDisplayUI
	v := reflect.ValueOf(human)
	for i := 0; i < v.NumField(); i++ {
		label := canvas.NewText(v.Type().Field(i).Name, color.White)
		var text *canvas.Text
		value := v.Field(i)
		log.Println(value)
		switch value.Kind() {
		case reflect.Int:
			text = canvas.NewText(fmt.Sprintf("%d", value.Int()), color.White)
		case reflect.Bool:
			if value.Bool() {
				text = canvas.NewText("Yes", color.White)
			} else {
				text = canvas.NewText("No", color.White)
			}
		default:
			text = canvas.NewText(value.String(), color.White)
		}
		pairs = append(pairs, humanDisplayUI{Label: label, Text: text})
	}
	return pairs
}
