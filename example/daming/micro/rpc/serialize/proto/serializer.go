package proto

import (
	"errors"
	"google.golang.org/protobuf/proto"
)

type Serializer struct {
}

func (s *Serializer) Code() uint8 {
	return 2 // proto使用2
}

func (s *Serializer) Encode(val any) ([]byte, error) {
	msg, ok := val.(proto.Message)
	if !ok {
		return nil, errors.New("micro: 必须是proto.Message")
	}
	return proto.Marshal(msg)
}

func (s *Serializer) Decode(data []byte, val any) error {
	msg, ok := val.(proto.Message)
	if !ok {
		return errors.New("micro: 必须是proto.Message")
	}
	return proto.Unmarshal(data, msg)
}
