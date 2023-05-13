package main

import "fmt"

type base struct {
	num int
}

//fprintf writes formatted text to the output stream you specify.

// printf is equivalent to writing fprintf(stdout, ...) and writes formatted text to wherever the standard output stream is currently pointing.

// sprintf writes formatted text to an array of char, as opposed to a stream.

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
} // base is embedded in container

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other structs
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
