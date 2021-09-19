# escape analysis

```shell script
go build -gcflags="-m" main.go
# command-line-arguments
./main.go:3:6: can inline main
./main.go:4:14: make([]int, 10240) escapes to heap
```

-gcflags: go compile flags