package common_test

import (
	"testing"

	"trachtenberg-math-system/common"
)

func TestLargeNumber(t *testing.T) {
	validLargeNums := []string{
		"1984903848193841",
		"1083401843814818341374187348173417384718327481738473287417",
		"9734817347134718371734137483641374138749813787413874817173741837482173471328748173748127348173847182374",
	}
	for _, l := range validLargeNums {
		ln, err := common.NewLargeNumber(l)
		if err != nil {
			t.Errorf("LargeNumber: %v failed to verify", ln)
			t.FailNow()
		}
	}
	invalidLargeNumArr := []string{
		"2048024802183948thoenuhaoeut0913248092",
		"(*)*&(*&*ohuhah23334213aoeuaoeu",
		"&&&&&&~~~~~~!!!!!!!!@############234323312333344",
	}
	for _, l := range invalidLargeNumArr {
		ln, err := common.NewLargeNumber(l)
		if err == nil {
			t.Errorf("LargeNumber: %v, should not be valid and an error should be thrown", ln)
			t.FailNow()
		}
	}
}

func TestVeryLargeNumber(t *testing.T) {
	largeNumStr := ""
	baseLargeNum := "123456789"
	for i := 0; i < 10000; i++ {
		largeNumStr += baseLargeNum
	}
	largeNum, err := common.NewLargeNumber(largeNumStr)
	if err != nil {
		t.Errorf("A very large number: %v, should not throw any error", largeNum)
		t.FailNow()
	}
}

func TestNewLargeNumber(t *testing.T) {
	ln, err := common.NewLargeNumber("120984102384")
	if err != nil {
		t.Errorf("Error: %v, should not have occurred for LargeNumber: %v", err, ln)
		t.FailNow()
	}
	ln, err = common.NewLargeNumber("ntaoehua083140912384")
	if err == nil {
		t.Errorf("LargeNumber: %v, should throw an error", ln)
		t.FailNow()
	}
}

func TestCurrentDigit(t *testing.T) {
	ln, err := common.NewLargeNumber("123456789009")
	if err != nil {
		t.Errorf("Creating new large number should not have failed")
		t.FailNow()
	}
	if ln.CurrentDigit() != 9 {
		t.Error("Current digit is supposed to be NINE")
		t.FailNow()
	}
	ln.MoveToNextIndex()
	if ln.CurrentDigit() != 0 {
		t.Error("Current digit is supposed to be ZERO")
		t.FailNow()
	}
}

func TestCurrentIndex(t *testing.T) {
	ln, err := common.NewLargeNumber("274987214134141")
	if err != nil {
		t.Error("New large number creation should not have failed")
		t.FailNow()
	}
	if ln.CurrentIndex() != len(ln.String())-1 {
		t.Error("Wrong initial current index")
		t.FailNow()
	}
}

func TestMoveToNextIndex(t *testing.T) {
	largeNumStrs := []string{"173497197", "87937149817", "48493025312"}
	largeNumArrs := [][]int{
		{1, 7, 3, 4, 9, 7, 1, 9, 7},
		{8, 7, 9, 3, 7, 1, 4, 9, 8, 1, 7},
		{4, 8, 4, 9, 3, 0, 2, 5, 3, 1, 2}}
	for k, lStr := range largeNumStrs {
		ln, err := common.NewLargeNumber(lStr)
		if err != nil {
			t.Error("Error should not be raised for a valid number")
			t.FailNow()
		}
		numAsArr := largeNumArrs[k]
		for ln.CurrentIndex() != -1 {
			if ln.CurrentDigit() != numAsArr[ln.CurrentIndex()] {
				t.Error("Digits don't match")
				t.FailNow()
			}
			ln.MoveToNextIndex()
		}
	}
}

func TestPreviousIndex(t *testing.T) {
	largeNumStrs := []string{"71937491274", "3498298790901234", "87283748274"}
	for _, lStr := range largeNumStrs {
		prevDigit := 0
		ln, err := common.NewLargeNumber(lStr)
		for ln.CurrentIndex() != -1 {
			if err != nil {
				t.Error("Error should not have raised for a valid number")
				t.FailNow()
			}
			if ln.PreviousDigit() != prevDigit {
				t.Error("Previous digits don't match")
				t.FailNow()
			}
			prevDigit = ln.CurrentDigit()
			ln.MoveToNextIndex()
		}
	}
}

func TestResetIndexTo(t *testing.T) {
	ln, err := common.NewLargeNumber("82374823748273")
	if err != nil {
		t.Error("Error should not be thrown for a valid large number")
		t.FailNow()
	}
	ln.MoveToNextIndex()
	ln.MoveToNextIndex()
	if ln.CurrentDigit() != 2 {
		t.Error("Current digit is supposed to be TWO")
		t.FailNow()
	}
	ln.ResetIndexTo(1)
	if ln.CurrentDigit() != 3 {
		t.Error("Current digit should be 3, but received: ", ln.CurrentDigit())
		t.FailNow()
	}
	for i := 0; i < 5; i++ {
		ln.MoveToNextIndex()
	}
	ln.ResetIndexTo(-1)
	if ln.CurrentDigit() != 4 {
		t.Error("Current digit should be 4, but received: ", ln.CurrentDigit())
		t.FailNow()
	}
	ln.ResetIndexTo(-1)
	if ln.CurrentDigit() != 8 {
		t.Error("Current digit should be 8, but received: ", ln.CurrentDigit())
		t.FailNow()
	}
	prevIndx := ln.CurrentIndex()
	ln.ResetIndexTo(10)
	if prevIndx != ln.CurrentIndex() {
		t.Error("Current index should not have changed")
		t.FailNow()
	}
}
