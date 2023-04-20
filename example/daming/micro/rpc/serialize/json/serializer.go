package json

import "encoding/json"

type Serializer struct {
}

func (s *Serializer) Code() uint8 {
	return 1 // json使用1
}

func (s *Serializer) Encode(val any) ([]byte, error) {
	return json.Marshal(val)
}

func (s *Serializer) Decode(data []byte, val any) error {
	return json.Unmarshal(data, val)
}
