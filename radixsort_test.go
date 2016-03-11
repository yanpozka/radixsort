package radixsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSliceDigit(t *testing.T) {
	nums := []int{6, 123, 29, 0, 9876}

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
			assert.Equal(t, result, digit(nums, ix, pos+1), "")
		}
	}
}

func TestAmountDigits(t *testing.T) {
	nums := []int{1234, 13, 9, 123456789}

	assert := assert.New(t)
	assert.Equal(amountDigits(nums, 0), 4)
	assert.Equal(amountDigits(nums, 1), 2)
	assert.Equal(amountDigits(nums, 2), 1)
	assert.Equal(amountDigits(nums, 3), 9)
}

func TestRadixSort(t *testing.T) {
	a := []int{89, 5, 1, 623, 47, 11}
	RadixSort(a)

	assert.Equal(t, []int{1, 5, 11, 47, 89, 623}, a, "Array A should be ordered")

	b := []int{1, 12, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	RadixSort(b)

	assert.Equal(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 12}, b, "Array b should be ordered")
}
