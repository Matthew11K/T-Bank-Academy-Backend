package benchmarks_test

import (
	"sync"
	"testing"
)

func addWithChannelChunked(data []int, numOfGoroutines int) int {
	sum := 0
	ch := make(chan int, numOfGoroutines)

	var wg sync.WaitGroup

	chunkSize := len(data) / numOfGoroutines
	remainder := len(data) % numOfGoroutines

	for i := 0; i < numOfGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == numOfGoroutines-1 {
			end += remainder
		}

		wg.Add(1)

		go func(chunk []int) {
			defer wg.Done()

			localSum := 0

			for _, v := range chunk {
				localSum += v
			}
			ch <- localSum
		}(data[start:end])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for partialSum := range ch {
		sum += partialSum
	}

	return sum
}

func addWithMutexChunked(data []int, numOfGoroutines int) int {
	sum := 0
	mu := sync.Mutex{}

	var wg sync.WaitGroup

	chunkSize := len(data) / numOfGoroutines
	remainder := len(data) % numOfGoroutines

	for i := 0; i < numOfGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == numOfGoroutines-1 {
			end += remainder
		}

		wg.Add(1)

		go func(chunk []int) {
			defer wg.Done()

			localSum := 0

			for _, v := range chunk {
				localSum += v
			}

			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(data[start:end])
	}

	wg.Wait()

	return sum
}

func addWithWaitGroup(data []int) int {
	sum := 0

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		localSum := 0

		for _, v := range data {
			localSum += v
		}

		sum += localSum
	}()

	wg.Wait()

	return sum
}

func BenchmarkAddWithChannel(b *testing.B) {
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}

	numOfGoroutines := 4

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = addWithChannelChunked(data, numOfGoroutines)
	}
}

func BenchmarkAddWithMutexChunked(b *testing.B) {
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}

	numOfGoroutines := 4

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = addWithMutexChunked(data, numOfGoroutines)
	}
}

func BenchmarkAddWithWaitGroup(b *testing.B) {
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = addWithWaitGroup(data)
	}
}
