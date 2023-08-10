package main

import "fmt"

// ReverseIterator represents a reverse iterator for a slice.
type ReverseIterator struct {
	slice []int
	index int
}

// NewReverseIterator creates a new reverse iterator for the given slice.
func NewReverseIterator(slice []int) *ReverseIterator {
	return &ReverseIterator{
		slice: slice,
		index: len(slice) - 1,
	}
}

// Next returns the next element in the reverse iteration.
func (ri *ReverseIterator) Next() (int, bool) {
	if ri.index >= 0 {
		value := ri.slice[ri.index]
		ri.index--
		return value, true
	}
	return 0, false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Iterate over the slice in the reverse direction using the custom ReverseIterator.
	for it := NewReverseIterator(nums); ; {
		if value, ok := it.Next(); ok {
			fmt.Println(value)
		} else {
			break
		}
	}

	for _, x := range nums {
		fmt.Println(x)
	}
}
