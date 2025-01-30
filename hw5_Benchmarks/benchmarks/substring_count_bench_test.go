package benchmarks_test

import (
	"math/rand"
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/pkg/utils"
)

func generateRandomString(size int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, size)

	for i := range s {
		//nolint:gosec // не требуется сильная рандомизация
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func BenchmarkCountSubstringsStringsCount(b *testing.B) {
	s := generateRandomString(10000)

	const substr = "test"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		utils.CountSubstringsStringsCount(s, substr)
	}
}

func BenchmarkCountSubstringsRegex(b *testing.B) {
	s := generateRandomString(10000)
	substr := "test"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = utils.CountSubstringsRegex(s, substr)
	}
}

func BenchmarkCountSubstringsManual(b *testing.B) {
	s := generateRandomString(10000)
	substr := "test"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		utils.CountSubstringsManual(s, substr)
	}
}
