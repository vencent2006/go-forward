# zero garbage

```shell script
go test -bench=. -benchmem

goos: darwin
goarch: amd64
pkg: go-examples/zero_garbage
BenchmarkZeroGarbage-8   	57015768	        18.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasGarbage-8    	58085676	        18.8 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	go-examples/zero_garbage	2.210s
```


结论：使用sync.Pool避免对象分配