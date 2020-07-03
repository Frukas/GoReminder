package main

import (
	"fyne.io/fyne"
)

var timeLabel fyne.CanvasObject

var ap = newAppWindow()

func main() {

	var con screenContainner
	con.init()

	ap.newContestWindowAndRun(con.setNewContainer())

}
