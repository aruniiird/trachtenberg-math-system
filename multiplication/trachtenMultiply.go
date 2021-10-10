package multiplication

import (
	"fmt"
	"strings"
	"trachtenberg-math-system/common"
)

var (
	unitsArr = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 2, 4, 6, 8, 0, 2, 4, 6, 8},
		{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		{0, 4, 8, 2, 6, 0, 4, 8, 2, 6},
		{0, 5, 0, 5, 0, 5, 0, 5, 0, 5},
		{0, 6, 2, 8, 4, 0, 6, 2, 8, 4},
		{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		{0, 8, 6, 4, 2, 0, 8, 6, 4, 2},
		{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}

	tensArr = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 2, 2, 2, 3, 3},
		{0, 0, 1, 1, 2, 2, 3, 3, 4, 4},
		{0, 0, 1, 1, 2, 3, 3, 4, 4, 5},
		{0, 0, 1, 2, 2, 3, 4, 4, 5, 6},
		{0, 0, 1, 2, 3, 4, 4, 5, 6, 7},
		{0, 0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
)

func Multiplication(multiplicant, multiplier *common.LargeNumber) *common.LargeNumber {
	return MultiplicationWithVerbosity(multiplicant, multiplier, false)
}

func MultiplicationWithVerbosity(multiplicant, multiplier *common.LargeNumber, verbose bool) *common.LargeNumber {
	var productStr string
	var carry int
	multiplicant.Reset()
	multiplier.Reset()
	if verbose {
		fmt.Println("Multiplicant: ", multiplicant)
		fmt.Println("Multiplier: ", multiplier)
	}
	for {
		if verbose {
			fmt.Println("------------------------")
		}
		var sum int
		for {
			mrD := multiplier.CurrentDigit()
			mtD := multiplicant.CurrentDigit()
			mtPrevD := multiplicant.PreviousDigit()

			u := unitsArr[mrD][mtD]
			t := tensArr[mrD][mtPrevD]
			sum += u + t

			if verbose {
				fmt.Println("MrD: ", mrD, ",  MtD: ", mtD, ", PrevMtD: ", mtPrevD, ", Unit: ", u, ", Ten: ", t, ", U+T: ", (u + t))
			}

			if move := multiplicant.MoveToPreviousIndex(); !move {
				break
			}
			if move := multiplier.MoveToNextIndex(); !move || multiplier.CurrentIndex() == -1 {
				break
			}
		}
		if verbose {
			fmt.Println("All U+T Sums: ", sum, ", PrevCarry: ", carry)
		}
		sum += carry
		carry = sum / 10
		productStr = fmt.Sprintf("%d%s", sum%10, productStr)
		multiplier.ResetIndexToBookmark()
		multiplicant.ResetIndexToBookmark()
		if moved := multiplicant.MoveToNextIndex(); !moved {
			multiplier.MoveToNextIndex()
			multiplier.SaveAsBookmark()
		} else {
			multiplicant.SaveAsBookmark()
		}
		if multiplier.CurrentIndex() == -1 {
			break
		}
	}
	productStr = strings.TrimLeft(productStr, "0")
	if productStr == "" {
		productStr = "0"
	}

	multiplicant.Reset()
	multiplier.Reset()

	productLn, _ := common.NewLargeNumber(productStr)
	return productLn
}
