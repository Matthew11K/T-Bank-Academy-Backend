package sorting_test

import (
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/pkg/sorting"

	"github.com/stretchr/testify/assert"
)

func TestStandardSort(t *testing.T) {
	sorter := sorting.NewSorter()
	data := []int{5, 3, 2, 4, 1}

	sorter.StandardSort(data)

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, data)
}

func TestBubbleSort(t *testing.T) {
	sorter := sorting.NewSorter()
	data := []int{5, 3, 2, 4, 1}

	sorter.BubbleSort(data)

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, data)
}
