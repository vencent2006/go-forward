package serialize

type Serializer interface {
	Code() uint8
	Encode(val any) ([]byte, error)
	// val 应该是一个结构体指针
	Decode(data []byte, val any) error
}
