# error包的使用

主要是针对go 1.13之后




* [Go 1.13 errors 基本用法](https://www.cnblogs.com/twoheads/p/12845404.html)
* [Go1.13 errors包用法](https://blog.csdn.net/ChamPly/article/details/100552448)

```go
	// New
	fmt.Println("-------New--------")
	err1 := errors.New("error1") 
	err2 := errors.New("error2")

	// Is
	fmt.Println("-------Is--------")
	fmt.Println(errors.Is(err1, err2)) //false
	fmt.Println(errors.Is(err1, errors.New("error1")))//false

	// As
	fmt.Println(errors.As(err1, &err2)) //true; As的话，值相同就可以了；Is必须类型也得相同
	fmt.Println("-------As--------")
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {// true
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
```