# big endian vs little endian


```go
// golang
a := uint32(0x01020304)
arr := make([]byte, 4)
binary.BigEndian.PutUint32(arr, a)
t.Log(arr) //[1 2 3 4]

binary.LittleEndian.PutUint32(arr, a)
t.Log(arr) //[4 3 2 1]
```