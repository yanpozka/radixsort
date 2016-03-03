package radixsort

import (
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {

	RadixSort(sort.IntSlice([]int{89, 5, 1, 23, 47, 11}))

}
