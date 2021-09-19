# string concat

```shell script
go test -bench=. 

goos: darwin
goarch: amd64
pkg: go-examples/string_concat
BenchmarkPlusConcat-8           22749913                50.5 ns/op
BenchmarkSprintf-8               5609557               214 ns/op
PASS
ok      go-examples/string_concat       2.632s

```


可以看出，sprintf的效率不如直接字符串相加的