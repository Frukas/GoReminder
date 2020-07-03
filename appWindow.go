package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

type appWindow struct {
	window fyne.App
}

func newAppWindow() *appWindow {
	ap := app.New()
	var cTheme custTheme
	fyne.CurrentApp().Settings().SetTheme(cTheme)
	fmt.Printf("%T", fyne.CurrentApp().Settings().Theme())
	return &appWindow{window: ap}

}

func (a *appWindow) newContestWindow(f fyne.CanvasObject) fyne.Window {

	w := a.window.NewWindow("Reminder Tool")
	w.Resize(fyne.NewSize(700, 400))
	w.SetContent(f)
	w.Show()
	return w
}

func (a *appWindow) newContestWindowAndRun(f fyne.CanvasObject) fyne.Window {

	w := a.window.NewWindow("Reminder Tool")
	w.Resize(fyne.NewSize(700, 400))
	w.SetContent(f)
	w.ShowAndRun()
	return w
}
