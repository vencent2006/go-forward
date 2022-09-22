package main

import "fmt"

func main() {
	m := make(map[string]int, 10)
	fmt.Println(m)
	hash := map[string]int{
		"1": 2,
		"3": 4,
		"5": 6,
	}
	fmt.Println(hash)
}

/*
(base) [17:03:34] 4_5   master ●✚  go tool compile -S -N -l main.go | grep "main.go:6"
        0x0021 00033 (main.go:6)        LEAQ    type.map[string]int(SB), AX
        0x0028 00040 (main.go:6)        MOVQ    AX, (SP)
        0x002c 00044 (main.go:6)        MOVQ    $10, 8(SP)
        0x0035 00053 (main.go:6)        MOVQ    $0, 16(SP)
        0x003e 00062 (main.go:6)        PCDATA  $1, $0
        0x003e 00062 (main.go:6)        NOP
        0x0040 00064 (main.go:6)        CALL    runtime.makemap(SB)
        0x0045 00069 (main.go:6)        MOVQ    24(SP), AX
        0x004a 00074 (main.go:6)        MOVQ    AX, "".m+48(SP)

(base) [17:04:33] 4_5   master ●✚  go tool compile -S -N -l main.go | grep "main.go:8"
        0x00d5 00213 (main.go:8)        CALL    runtime.makemap_small(SB)
        0x00da 00218 (main.go:8)        MOVQ    (SP), AX
        0x00de 00222 (main.go:8)        MOVQ    AX, "".hash+64(SP)

*/
