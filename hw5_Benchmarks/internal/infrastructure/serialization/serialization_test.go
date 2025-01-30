package serialization_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_5-go-Matthew11K/internal/infrastructure/serialization"
)

func TestJSONSerialization(t *testing.T) {
	serializer := serialization.NewSerializer()
	data := domain.Data{ID: 1, Name: "Test"}

	serialized, err := serializer.SerializeJSON(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, serialized)

	var deserialized domain.Data
	err = serializer.DeserializeJSON(serialized, &deserialized)
	assert.NoError(t, err)
	assert.Equal(t, data, deserialized)
}

func TestGobSerialization(t *testing.T) {
	serializer := serialization.NewSerializer()
	data := domain.Data{ID: 2, Name: "GobTest"}

	serialized, err := serializer.SerializeGob(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, serialized)

	var deserialized domain.Data
	err = serializer.DeserializeGob(serialized, &deserialized)
	assert.NoError(t, err)
	assert.Equal(t, data, deserialized)
}

func TestMsgPackSerialization(t *testing.T) {
	serializer := serialization.NewSerializer()
	data := domain.Data{ID: 3, Name: "MsgPackTest"}

	serialized, err := serializer.SerializeMsgPack(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, serialized)

	var deserialized domain.Data
	err = serializer.DeserializeMsgPack(serialized, &deserialized)
	assert.NoError(t, err)
	assert.Equal(t, data, deserialized)
}
