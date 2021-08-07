package main

import (
	"fmt"
	"strconv"
	"strings"
)

//TimeFormat is struct of time for the format hh:mm:ss
type TimeFormat struct {
	hour   int
	min    int
	sec    int
	stHour int
	stMin  int
	stsec  int
}

func (t *TimeFormat) toString() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.hour, t.min, t.sec)
}

func (t *TimeFormat) substractSec() {
	if t.sec > 0 {
		t.sec--
	} else {
		if t.min > 0 {
			t.min--
			t.sec = 59
		} else {
			if t.hour > 0 {
				t.hour--
				t.min = 59
				t.sec = 59
			}
		}
	}
}

func (t *TimeFormat) isZeroTime() bool {
	if t.hour <= 0 && t.min <= 0 && t.sec <= 0 {
		return true
	}
	return false
}

func (t *TimeFormat) setNewTime(time string) {
	s := strings.Split(time, ":")
	var err error
	t.hour, err = strconv.Atoi(s[0])
	errorHandler(err, s[0])
	t.stHour = t.hour
	t.min, err = strconv.Atoi(s[1])
	errorHandler(err, s[1])
	t.stMin = t.min
	if len(s) > 2 {
		t.sec, _ = strconv.Atoi(s[2])
		errorHandler(err, s[2])
		t.stsec = t.sec
	} else {
		t.sec = 0
		t.stsec = t.sec
	}
}

func (t *TimeFormat) reset() {
	t.hour = t.stHour
	t.min = t.stMin
	t.sec = t.stsec
}

func (t *TimeFormat) resetToZero() {
	t.hour = 0
	t.min = 0
	t.sec = 0
	t.stHour = 0
	t.stMin = 0
	t.stsec = 0
}

func errorHandler(e error, s string) {
	if e != nil {
		fmt.Printf("There was error converting the value %v to integer. Error message %v", s, e.Error())
	}
}

func (t *TimeFormat) checkStringTonumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("It was not possible to convert the String ", s, " to an integer")
		return false
	}
	return true
}

func (t *TimeFormat) formatingMin(s string) string {
	min, err := strconv.Atoi(s)
	errorHandler(err, s)

	fhour := min / 60
	fmin := min % 60
	if min < 0 {
		return "00:00:00"
	}

	return fmt.Sprintf("%02d:%02d:00", fhour, fmin)
}
