package benchmarks_test

import (
	"math/rand"
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/pkg/sorting"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		//nolint:gosec // не требуется сильная рандомизация
		slice[i] = rand.Intn(size)
	}

	return slice
}

func BenchmarkStandardSort100(b *testing.B) {
	sorter := sorting.NewSorter()
	data := generateRandomSlice(100)

	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(data))

		b.StopTimer()
		copy(tmp, data)
		b.StartTimer()

		sorter.StandardSort(tmp)
	}
}

func BenchmarkBubbleSort100(b *testing.B) {
	sorter := sorting.NewSorter()
	data := generateRandomSlice(100)

	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(data))

		b.StopTimer()
		copy(tmp, data)
		b.StartTimer()

		sorter.BubbleSort(tmp)
	}
}

func BenchmarkStandardSort1000(b *testing.B) {
	sorter := sorting.NewSorter()
	data := generateRandomSlice(1000)

	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(data))

		b.StopTimer()
		copy(tmp, data)
		b.StartTimer()

		sorter.StandardSort(tmp)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	sorter := sorting.NewSorter()
	data := generateRandomSlice(1000)

	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(data))

		b.StopTimer()
		copy(tmp, data)
		b.StartTimer()

		sorter.BubbleSort(tmp)
	}
}
