/**
 * @Author: vincent
 * @Description:
 * @File:  token_test
 * @Version: 1.0.0
 * @Date: 2021/9/23 15:17
 */

package ch01

import (
	"fmt"
	"go/scanner"
	"go/token"
	"testing"
)

func TestToken(t *testing.T) {
	src := []byte("cos(x) + 2i*sin(x) // Euler")

	// 初始化scanner
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	// 扫描
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
