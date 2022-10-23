package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/fyne-io/terminal"
)

func main() {
	CreateTerminal()
}


func CreateTerminal(){
	myApp := app.New()
	myWindow := myApp.NewWindow("Pop Terminal")
	t := terminal.New()
	myWindow.SetContent(t)

	go func() {
		_ = t.RunLocalShell()
		//a.Quit()
	}()
	myWindow.ShowAndRun()
}
