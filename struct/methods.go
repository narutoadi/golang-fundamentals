package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) is_square() bool {
	if r.width == r.height {
		return true
	} else {
		return false
	}
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	sq := rect{width: 10, height: 10}

	fmt.Println(sq, "is it square?", sq.is_square())
}
