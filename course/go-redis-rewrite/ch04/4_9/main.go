package main

import "fmt"

type Car interface {
	Drive()
}

type Truck struct {
	Model string
}

func (t Truck) Drive() {
	fmt.Println(t.Model)
}

type TrafficTool interface {
	Drive()
	Run()
}

func main() {
	var c Car = Truck{}
	t := c.(Truck)
	tt, err := c.(TrafficTool)
	if err != true {
		fmt.Println(err)
	}

	switch c.(type) {
	case TrafficTool:
		fmt.Println("tt")
	}
	fmt.Println(t)
	fmt.Println(tt)
}
