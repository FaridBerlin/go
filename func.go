package main

import "fmt"

func callfunc(callable func(int) int) int {
	return callable(5)
}

func main() {

	value := callfunc(func(x int) int {
		return x * x
	})

	fmt.Println(value)
	
}