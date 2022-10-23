package main

import (
	//"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

type User struct{
	Name string
	ID string
}

func main() {
	user := User{
		Name:"Mojjo Jojo",
		ID: "okjiuhbgvycrtybu",
	}
	myApp := app.New()
	w := myApp.NewWindow("WHEAGLE")

	//gradient := canvas.NewHorizontalGradient(color.Black, color.Transparent)
	//gradient := canvas.NewRadialGradient(color.Black, color.Transparent)
	//w.SetContent(gradient)
	top := widget.NewLabel(user.Name)
	bottom := widget.NewButton("Hi",func(){
		top.SetText("Welcome:")
	})
	w.SetContent(container.NewVSplit(
		top,
		bottom,
		))

	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}


/*menuItem := &fyne.Menu{
	Label: "Minions",
	Items: nil,
}
menu := fyne.NewMainMenu(menuItem)
w.SetMainMenu(menu)*/
