package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)

	s1 := make([]int, 10)
	fmt.Println(s1)
}

/*
查看汇编语言
go build -gcflags -S main.go
FUNCDATA        $3, "".main.stkobj(SB)
        0x0021 00033  (main.go:6)   LEAQ    type.[3]int(SB), AX
        0x0028 00040  (main.go:6)   MOVQ    AX, (SP)
        0x002c 00044  (main.go:6)   PCDATA  $1, $0
        0x002c 00044  (main.go:6)   CALL    runtime.newobject(SB)  // runtime.slice
        0x0031 00049  (main.go:6)   MOVQ    8(SP), AX
        0x0036 00054  (main.go:6)   MOVQ    $1, (AX)
        0x003d 00061  (main.go:6)   MOVQ    $2, 8(AX)
        0x0045 00069  (main.go:6)   MOVQ    $3, 16(AX)
        0x004d 00077  (main.go:7)   MOVQ    AX, (SP)
        0x0051 00081  (main.go:7)
*/

/*
容量扩容的逻辑
runtime/slice.go growslice
*/
