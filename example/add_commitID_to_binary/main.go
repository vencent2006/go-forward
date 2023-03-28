package main

import "fmt"

var (
	Tag      = "v0.0.0"
	CommitID = ""
	Branch   = ""
	DATE     = ""
)

func main() {
	fmt.Println("tag:", Tag, "branch:", Branch, "commitID:", CommitID, "DATE:", DATE)
}
