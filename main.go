package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func insbutttonf(HowManySent *int, a fyne.App, ideabox fyne.App) {
	PassedregulatednumbersSentErrorWindow := a.NewWindow("Error")
	if *HowManySent <= 0 {

		PassedregulatednumbersSentErrorWindowError := widget.NewLabel("sorry but you passed you 3 messages a month limit")
		PassedregulatednumbersSentErrorWindow.SetContent(PassedregulatednumbersSentErrorWindowError)
		PassedregulatednumbersSentErrorWindow.Show()
	} else {

		fmt.Println("button Pressed", *HowManySent, "requests remain")
		fmt.Println(ideabox)
		*HowManySent--
	}
}
func main() {
	// var RecentSent([3]string)
	var HowManySent int = 3
	a := app.New()
	wi := a.NewWindow("Melkey's Idea")
	wi.Resize(fyne.NewSize(1200, 400))
	question := widget.NewLabel("Insert your Idea for Melkey")
	logo := canvas.NewImageFromFile("Icon.png")
	logo.FillMode = canvas.ImageFillStretch
	ideabox := widget.NewEntry()
	ideabox.SetPlaceHolder("Your idea ")
	insbutton := widget.NewButtonWithIcon("Enter", fyne.Resource(logo.Resource), func() { insbutttonf(&HowManySent, a) })
	ideaboxs := container.NewGridWithColumns(2, ideabox, insbutton)
	rectidea := container.NewGridWithRows(2, question, ideaboxs)
	wi.SetContent(rectidea)
	wi.ShowAndRun()
}
