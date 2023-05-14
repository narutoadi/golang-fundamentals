package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i <= 3; i++ {
		fmt.Println("in loop", i, &i)
		func() {
			fmt.Println("inside closure function", i, &i)
		}()
	}

	for i := 0; i <= 3; i++ {
		fmt.Println("in loop", i, &i)
		i := i
		go func() {
			fmt.Println("inside goroutine closure function", i, &i)
		}()
	}
	time.Sleep(time.Second)

}
