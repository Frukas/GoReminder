package main

import "testing"

func TestToString(t *testing.T) {
	teForm := TimeFormat{hour: 4, min: 3, sec: 02}
	correctResult := "04:03:02"

	if teForm.toString() != correctResult {
		t.Error("Expected: ", correctResult, "got: ", teForm.toString())
	}
}

func TestSubstractSec(t *testing.T) {
	teForm := []TimeFormat{
		{4, 3, 2, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
	}
	correctResult := []string{
		"04:03:01",
		"00:00:00",
		"00:59:59",
		"00:00:59",
	}

	for i, v := range teForm {
		v.substractSec()
		if v.toString() != correctResult[i] {
			t.Error("Expected: ", correctResult[i], "got:", v.toString(), "of the test:", teForm[i].toString())
		}
	}

}

func TestIsZeroTime(t *testing.T) {
	teForm := []TimeFormat{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{0, -1, 0, 0, 0, 0},
		{0, 0, -1, 0, 0, 0},
	}
	correctResult := []bool{
		true,
		false,
		false,
		false,
		false,
		true,
		true,
	}

	for i, v := range teForm {
		if v.isZeroTime() != correctResult[i] {
			t.Error("Expected ", correctResult[i], "got ", v.isZeroTime(), "of the test ", teForm[i].toString())
		}
	}
}

func TestFormatingMin(t *testing.T) {
	var tm TimeFormat
	testmap := map[string]string{
		"124": "02:04:00",
		"34":  "00:34:00",
		"61":  "01:01:00",
		"300": "05:00:00",
		"0":   "00:00:00",
		"-1":  "00:00:00",
	}

	for m, v := range testmap {
		if tm.formatingMin(m) != v {
			t.Error("for the Value ", m, "expected result ", v, "but got ", tm.formatingMin(m))
		}
	}
}

func TestCheckStringTonumber(t *testing.T) {
	var tm TimeFormat
	testmap := map[string]bool{
		"2":   true,
		"34":  true,
		"6:1": false,
		"3i0": false,
		"abc": false,
		"-1":  true,
	}

	for m, v := range testmap {
		if tm.checkStringTonumber(m) != v {
			t.Error("for the Value ", m, "expected result ", v, "but got ", tm.formatingMin(m))
		}
	}
}
