package message

type Response struct {
	HeadLength uint32
	BodyLength uint32
	RequestID  uint32
	Version    uint8
	Compresser uint8
	Serializer uint8

	Error []byte

	Data []byte
}
