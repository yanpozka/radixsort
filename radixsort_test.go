package radixsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSliceDigit(t *testing.T) {
	nums := IntRadixSlice([]int{6, 123, 29, 0, 9876})

	expected := [][]int{
		//1  2  3  4
		{6, 0, 0, 0},
		{3, 2, 1, 0},
		{9, 2, 0, 0},
		{0, 0, 0, 0},
		{6, 7, 8, 9},
	}

	for ix := range nums {
		for pos, result := range expected[ix] {
			assert.Equal(t, result, nums.Digit(ix, pos+1), "")
		}
	}
}

func TestRadixSort(t *testing.T) {

	RadixSort(IntRadixSlice([]int{89, 5, 1, 23, 47, 11}))

}
