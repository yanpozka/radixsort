package radixsort

import (
	"container/list"
	"fmt"
	"math"
	"sort"
)

func digit(n, pos int) int {
	if pos <= 0 || pos > 70 {
		return -1
	}
	return (n / int(math.Pow10(pos-1))) % 10
}

func amountDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

//
func RadixSort(data []int) {
	if sort.IntsAreSorted(data) {
		return
	}

	var max_digit int
	origin := list.New()

	for ix, size := 0, len(data); ix < size; ix++ {
		if ad := amountDigits(data[ix]); ad > max_digit {
			max_digit = ad
		}
		origin.PushBack(data[ix])
	}

	r := radixSort(origin, max_digit)

	for ix, elem := 0, r.Front(); elem != nil; ix, elem = ix+1, elem.Next() {
		data[ix] = elem.Value.(int)
	}
}

func radixSort(data *list.List, position int) *list.List {

	if position == 0 || data.Len() <= 1 {
		return data
	}

	var bucket [10]*list.List

	for elem := data.Front(); elem != nil; elem = elem.Next() {
		d := digit(elem.Value.(int), position)

		if bucket[d] == nil {
			bucket[d] = list.New()
		}

		bucket[d].PushBack(elem.Value)
	}

	output := make(chan *capsule, 10)
	var count int

	for ix, queue := range bucket {
		if queue == nil {
			continue
		}
		count++

		go func(i int, q *list.List, out chan *capsule) {
			l := radixSort(q, position-1)

			out <- &capsule{
				index: i,
				list:  l,
			}
		}(ix, queue, output)
	}

	var all [10]*list.List

	for ; count > 0; count-- {
		cr := <-output
		all[cr.index] = cr.list
	}

	result := list.New()
	for _, l := range all {
		if l == nil {
			continue
		}
		result.PushBackList(l)
	}

	return result
}

func printList(l *list.List) {
	if l == nil {
		fmt.Println("[Empty]")
		return
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

type capsule struct {
	index int
	list  *list.List
}
