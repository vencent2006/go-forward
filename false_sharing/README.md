# false sharing


```shell script
go test -bench=. -benchmem

goos: darwin
goarch: amd64
pkg: go-examples/false_sharing
BenchmarkNoPad-8               3         394019818 ns/op            5357 B/op         15 allocs/op
BenchmarkWithPad-8             4         265191220 ns/op            3032 B/op          8 allocs/op
PASS
ok      go-examples/false_sharing       3.615s

```

结论：withPad反倒在分配内存的次数和内存大小都是效率更高