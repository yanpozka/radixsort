package radixsort

import (
	"container/list"
	"math"
)

func digit(s []int, i, pos int) int {
	if pos <= 0 || pos > 70 {
		return -1
	}
	return (s[i] / int(math.Pow10(pos-1))) % 10
}

func amountDigits(s []int, i int) int {
	return int(math.Log10(float64(s[i]))) + 1
}

//
func RadixSort(data []int) {
	var max_digit int

	for ix, size := 0, len(data); ix < size; ix++ {
		if ad := amountDigits(data, ix); ad > max_digit {
			max_digit = ad
		}
	}

	radixSort(data, 0, len(data), max_digit)
}

func radixSort(data []int, start, end, position int) {
	size := end - start

	if position == 0 || size <= 1 {
		return
	}

	var bucket [10]*list.List

	for ix := 0; ix < size; ix++ {
		d := digit(data, ix, position)

		if bucket[d] == nil {
			bucket[d] = list.New()
		}

		bucket[d].PushBack(data[ix])
	}

	prev, count := start, start
	for _, q := range bucket {
		if q == nil {
			continue
		}

		for elem := q.Front(); elem != nil; elem = elem.Next() {
			data[count] = elem.Value.(int)
			count++ // !!
		}

		radixSort(data, prev, count, position-1)

		prev = count
	}
}
