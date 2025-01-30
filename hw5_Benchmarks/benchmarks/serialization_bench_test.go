package benchmarks_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/internal/infrastructure/serialization"
)

//nolint:gosec // не требуется сильная рандомизация
func generateTestData(size int) domain.Data {
	name := make([]byte, size)
	for i := range name {
		name[i] = byte(rand.Intn(26) + 65) // A-Z
	}

	tags := make(map[string]string)
	tagCount := rand.Intn(5) + 5

	for i := 0; i < tagCount; i++ {
		key := fmt.Sprintf("tag_%d", i)
		value := fmt.Sprintf("value_%d", rand.Intn(100))
		tags[key] = value
	}

	return domain.Data{
		ID:       rand.Int(),
		Name:     string(name),
		Tags:     tags,
		IsActive: rand.Float64() > 0.5,
		Score:    rand.Float64(),
	}
}

func BenchmarkSerializeJSON1KB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeJSON(data)
	}
}

func BenchmarkSerializeJSON1MB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(1024 * 1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeJSON(data)
	}
}

func BenchmarkSerializeJSON10MB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(10 * 1024 * 1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeJSON(data)
	}
}

func BenchmarkSerializeGob1KB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeGob(data)
	}
}

func BenchmarkSerializeGob1MB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(1024 * 1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeGob(data)
	}
}

func BenchmarkSerializeGob10MB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(10 * 1024 * 1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeGob(data)
	}
}

func BenchmarkSerializeMsgPack1KB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeMsgPack(data)
	}
}

func BenchmarkSerializeMsgPack1MB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(1024 * 1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeMsgPack(data)
	}
}

func BenchmarkSerializeMsgPack10MB(b *testing.B) {
	serializer := serialization.NewSerializer()
	data := generateTestData(10 * 1024 * 1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = serializer.SerializeMsgPack(data)
	}
}
