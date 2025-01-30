package sorting

import (
	"sort"
)

type Sorter interface {
	StandardSort(data []int)
	BubbleSort(data []int)
}

type DefaultSorter struct{}

func NewSorter() Sorter {
	return &DefaultSorter{}
}

func (s *DefaultSorter) StandardSort(data []int) {
	sort.Ints(data)
}

func (s *DefaultSorter) BubbleSort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
