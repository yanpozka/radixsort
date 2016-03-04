package radixsort

import (
	"math"
	"sort"
)

type RadixInterface interface {
	sort.Interface
	Digit(index, pos int) int
	AmountDigits(index int) int
}

type IntRadixSlice []int

func (s IntRadixSlice) Len() int           { return len(s) }
func (s IntRadixSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntRadixSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s IntRadixSlice) Digit(i, pos int) int {
	if pos <= 0 || pos > 70 {
		return -1
	}
	return (s[i] / int(math.Pow10(pos-1))) % 10
}

func (s IntRadixSlice) AmountDigits(i int) int {
	return int(math.Log10(float64(s[i]))) + 1
}

//
func RadixSort(data RadixInterface) {
	var max_digit int

	for ix, size := 0, data.Len(); ix < size; ix++ {
		if ad := data.AmountDigits(ix); ad > max_digit {
			max_digit = ad
		}
	}

	radixSort(data, max_digit)
}

func radixSort(data RadixInterface, digit int) {

}
