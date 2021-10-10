package common

import "errors"

type LargeNumber struct {
	largeNumStr    string
	currentIndx    int
	bookmarkedIndx int
}

func NewLargeNumber(lnStr string) (*LargeNumber, error) {
	if lnStr == "" {
		lnStr = "0"
	}
	ln := &LargeNumber{largeNumStr: lnStr, currentIndx: len(lnStr) - 1, bookmarkedIndx: len(lnStr) - 1}
	if !ln.Verify() {
		return nil, errors.New("not a valid number error")
	}
	return ln, nil
}

func (ln *LargeNumber) String() string {
	return ln.largeNumStr
}

func (ln *LargeNumber) Length() int {
	return len(ln.largeNumStr)
}

func (ln *LargeNumber) Verify() bool {
	valid := true
	for k := range ln.largeNumStr {
		vInt := ln.largeNumStr[k] - '0'
		if vInt > 9 {
			valid = false
			break
		}
	}
	return valid
}

func (ln *LargeNumber) CurrentIndex() int {
	return ln.currentIndx
}

func (ln *LargeNumber) MoveToNextIndex() bool {
	if ln.currentIndx == -1 {
		return false
	}
	ln.currentIndx -= 1
	return true
}

func (ln *LargeNumber) MoveToPreviousIndex() bool {
	if ln.currentIndx == ln.Length()-1 {
		return false
	}
	ln.currentIndx += 1
	return true
}

func (ln *LargeNumber) Reset() {
	ln.ResetIndex()
	ln.SaveAsBookmark()
}

func (ln *LargeNumber) ResetIndex() {
	ln.currentIndx = ln.Length() - 1
}

func (ln *LargeNumber) ResetIndexTo(newIndx int) {
	var realIndx int
	if newIndx < 0 {
		realIndx = ln.currentIndx - newIndx
	} else {
		realIndx = ln.Length() - newIndx
	}
	if realIndx >= ln.Length() {
		realIndx = ln.Length() - 1
	}
	// you can always reset index backwards
	if realIndx < ln.currentIndx {
		return
	}
	ln.currentIndx = realIndx
}

func (ln *LargeNumber) ResetIndexToBookmark() {
	ln.currentIndx = ln.bookmarkedIndx
}

func (ln *LargeNumber) SaveAsBookmark() {
	ln.bookmarkedIndx = ln.currentIndx
}

func (ln *LargeNumber) digitAt(indx int) int {
	if indx < 0 {
		return 0
	}
	runeArr := []rune(ln.largeNumStr)
	if indx >= len(runeArr) {
		return 0
	}
	return int(runeArr[indx] - '0')
}

func (ln *LargeNumber) CurrentDigit() int {
	return ln.digitAt(ln.currentIndx)
}

func (ln *LargeNumber) PreviousDigit() int {
	prevIndx := ln.CurrentIndex() + 1
	return ln.digitAt(prevIndx)
}
