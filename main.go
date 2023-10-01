package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func insbutttonf(HowManySent int) {
	HowManySent ++
	fmt.Print("button Pressed", HowManySent, "requests remain")
	

}
func main() {
	// var RecentSent([3]string)
	 var HowManySent int
	a := app.New()
	wi := a.NewWindow("Melkey's Idea")
	wi.Resize(fyne.NewSize(1200, 400))
	question := widget.NewLabel("Insert your Idea for Melkey")
	logo := canvas.NewImageFromFile("Icon.png")
	logo.FillMode = canvas.ImageFillStretch
	ideabox := widget.NewEntry()
	ideabox.SetPlaceHolder("Your idea ")
	insbutton := widget.NewButtonWithIcon("Enter", fyne.Resource(logo.Resource), insbutttonf)
	ideaboxs := container.NewGridWithColumns(2, ideabox, insbutton)
	rectidea := container.NewGridWithRows(2, question, ideaboxs)
	wi.SetContent(rectidea)
	wi.ShowAndRun()
}
