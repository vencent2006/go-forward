package user

// 用来给reflect包的fields_test.go，测试User是外来包的情况下的case，对应TestIterateFields_FromAnotherPkg
type User struct {
	Name string
	age  int
}
