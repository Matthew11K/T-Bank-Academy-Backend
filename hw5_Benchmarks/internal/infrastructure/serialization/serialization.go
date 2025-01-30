package serialization

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/vmihailenco/msgpack/v5"
)

type Serializer interface {
	SerializeJSON(data interface{}) ([]byte, error)
	DeserializeJSON(data []byte, v interface{}) error
	SerializeGob(data interface{}) ([]byte, error)
	DeserializeGob(data []byte, v interface{}) error
	SerializeMsgPack(data interface{}) ([]byte, error)
	DeserializeMsgPack(data []byte, v interface{}) error
}

type DefaultSerializer struct{}

func NewSerializer() Serializer {
	return &DefaultSerializer{}
}

func (s *DefaultSerializer) SerializeJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (s *DefaultSerializer) DeserializeJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (s *DefaultSerializer) SerializeGob(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)

	return buf.Bytes(), err
}

func (s *DefaultSerializer) DeserializeGob(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	return decoder.Decode(v)
}

func (s *DefaultSerializer) SerializeMsgPack(data interface{}) ([]byte, error) {
	return msgpack.Marshal(data)
}

func (s *DefaultSerializer) DeserializeMsgPack(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
