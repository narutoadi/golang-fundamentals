// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

package main

import (
	"fmt"
	"reflect"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func make_line(chars ...string) {
	fmt.Print(chars[0], "")
	for i := 1; i < len(chars); i++ {
		fmt.Print("====", chars[i])
	}
	fmt.Println("")
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
	str := []string{"hello", "hello", "hello"}
	make_line(str...)

	char := "a"
	fmt.Print(char, reflect.TypeOf(char))
}
