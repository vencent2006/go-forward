package main

import (
	"flag"
	"log"
)

func main() {
	parseName()
	parseName2()
}

func parseName() {
	var name string
	flag.StringVar(&name, "name", "", "帮助信息")
	flag.Parse()
	log.Printf("name:%s", name)
}
func parseName2() {
	var name string
	flag.StringVar(&name, "name", "", "帮助信息")
	flag.Parse()
	log.Printf("name:%s", name)
}

func case1() {
	parseName()
	parseName()
}
func case2() {
	parseName()
	parseName2()
}
func case3() {
	var name string
	flag.StringVar(&name, "name", "", "帮助信息")
}
