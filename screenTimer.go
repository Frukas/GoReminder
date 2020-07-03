package main

import (
	"fmt"
	"image/color"
	"path/filepath"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type screenContainner struct {
	//timerCount *widget.Label
	timerCount        *canvas.Text
	checkBox          *widget.SelectEntry
	textBox           *widget.Entry
	topLabel          *canvas.Text
	startButton       *widget.Button
	timer             *TimeFormat
	stopButton        *widget.Button
	custTime          *widget.Entry
	plusWindowsButton *widget.Button
}

func (s *screenContainner) init() {

	s.timerCount = canvas.NewText("00:00:00", color.White)
	s.timerCount.TextSize = 20
	s.timerCount.Alignment = fyne.TextAlignCenter
	s.checkBox = s.createCheckBox()
	s.textBox = widget.NewMultiLineEntry()

	s.textBox.PlaceHolder = "Write Here"
	s.topLabel = canvas.NewText("Reminder", color.White)
	s.topLabel.Alignment = fyne.TextAlignCenter
	s.topLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Monospace: false}
	s.startButton, s.stopButton = s.createStartStopButton()
	s.timer = &TimeFormat{}
	s.custTime = widget.NewEntry()
	s.custTime.PlaceHolder = "minutes"
	s.plusWindowsButton = s.CreatePlusWindowsButton()
}

func (s *screenContainner) setNewContainer() fyne.CanvasObject {

	left := widget.NewVBox(s.timerCount, s.checkBox, s.custTime, s.startButton, s.stopButton, s.plusWindowsButton)
	top := s.topLabel
	right := s.textBox
	horizontalContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(top, nil, left, nil), top, left, right)

	return horizontalContainer
}

func (s *screenContainner) createCheckBox() *widget.SelectEntry {

	strngList := []string{"00:05:00", "00:10:00", "00:15:00", "00:30:00", "01:00:00", "03:00:00"}
	selectItens := widget.NewSelectEntry(strngList)
	selectItens.SetPlaceHolder("00:00:00")
	selectItens.OnChanged = func(input string) {
		if s.custTime.Text == "" {
			s.timer.setNewTime(input)
		}
	}

	return selectItens
}

func (s *screenContainner) setTimerCount(count string) {
	//s.timerCount.SetText(count)
	s.timerCount.Text = count
	s.timerCount.Refresh()

}

func (s *screenContainner) createStartStopButton() (*widget.Button, *widget.Button) {
	tickSecond := time.NewTicker(time.Second)
	flag := false
	stbutton := widget.NewButton("Start", func() {
		//fmt.Println(s.timer.toString())
		//s.setTimerCount(s.timer.toString())
		if s.custTime.Text != "" {
			if s.timer.checkStringTonumber(s.custTime.Text) {
				s.timer.setNewTime(s.timer.formatingMin(s.custTime.Text))
			} else {
				s.custTime.Text = ""
			}
		}
		go func() {

			for range tickSecond.C {
				s.setTimerCount(s.timer.toString())
				//s.timer.substractSec()
				if s.timer.isZeroTime() {
					s.setTimerCount(s.timer.toString())
					go s.alertScreenShow()
					s.timer.reset()
					return
				} else if flag {
					flag = false
					return
				}
				//s.setTimerCount(s.timer.toString())
				s.timer.substractSec()
			}
		}()

	})

	spButton := widget.NewButton("Stop", func() {
		flag = true

		s.timer.reset()
		s.setTimerCount("00:00:00")
		s.custTime.Text = ""
		s.custTime.Refresh()
		s.custTime.Refresh()
	})
	return stbutton, spButton
}

func (s *screenContainner) playSound(t *time.Ticker) {
	files, err := filepath.Glob("*.wav")

	if err != nil {
		panic(err)
	}
	for _, file := range files {
		SndPlaySoundW(file, SND_SYNC)
		for range t.C {
			SndPlaySoundW(file, SND_SYNC)
			fmt.Println()
		}

	}
}

func (s *screenContainner) alertScreenShow() {
	var alertBox fyne.Window
	tickSecond := time.NewTicker(3 * time.Second)

	title := canvas.NewText("Alert", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 30

	text := canvas.NewText(s.textBox.Text, color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 20
	buttonOK := widget.NewButton("ok", func() {
		alertBox.Close()
		tickSecond.Stop()

	})
	alertBox = ap.newContestWindow(fyne.NewContainerWithLayout(layout.NewBorderLayout(title, buttonOK, nil, nil), title, text, buttonOK))

	go s.playSound(tickSecond)

	alertBox.SetOnClosed(func() {
		alertBox.Close()
		tickSecond.Stop()
	})
}

func (s *screenContainner) addPlusWindows() {
	var con screenContainner
	con.init()

	ap.newContestWindow(con.setNewContainer())
}

func (s *screenContainner) CreatePlusWindowsButton() *widget.Button {
	res, err := fyne.LoadResourceFromPath("plus.png")

	if err != nil {
		fmt.Println("It was not possible to load the image. Error: ", err)
	}
	return widget.NewButtonWithIcon("", res, func() {
		go s.addPlusWindows()
	})
}
