package multiplication_test

import (
	"testing"
	"trachtenberg-math-system/common"
	"trachtenberg-math-system/multiplication"
)

func TestMultiplication(t *testing.T) {
	anonyData := []struct {
		m1   string
		m2   string
		prod string
	}{
		{m1: "2", m2: "3", prod: "6"},
		{m1: "7", m2: "3", prod: "21"},
		{m1: "8", m2: "5", prod: "40"},
		{m1: "0", m2: "919389483984", prod: "0"},
		{m1: "0", m2: "000000000000", prod: "0"},
		{m1: "919389483984", m2: "0", prod: "0"},
		{m1: "919389483984", m2: "919389483984", prod: "845277023260365792512256"},
		{m1: "2", m2: "30", prod: "60"},
		{m1: "1", m2: "1000000000000", prod: "1000000000000"},
		{m1: "736478163476341637647136910916397164916340132468734619012634786234", m2: "6234871239746991643976176723647163276734687",
			prod: "4591846520160325767592041428763310361755142362167742991346918232190028110957928210201597475871290755777898758"},
	}
	for _, testData := range anonyData {
		ln1, err := common.NewLargeNumber(testData.m1)
		if err != nil {
			t.Error("Error should not be thrown for a valid large number")
			t.FailNow()
		}
		ln2, err := common.NewLargeNumber(testData.m2)
		if err != nil {
			t.Error("Error should not be thrown for a valid large number")
			t.FailNow()
		}
		productLn := multiplication.MultiplicationWithVerbosity(ln1, ln2, true)
		if productLn.String() != testData.prod {
			t.Error("Products don't match TrachtProd: ", productLn, ", ActualProd: ", testData.prod)
			t.FailNow()
		}
	}
}
