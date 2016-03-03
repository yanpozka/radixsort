package radixsort

import (
	"fmt"
	"math"
	"sort"
)

type RadixInterface interface {
	sort.Interface
	Digit(pos int) int
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

//
func RadixSort(data sort.Interface) {
	fmt.Println(data)
}
