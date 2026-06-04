package main

import "fmt"

type person struct {
	name string
	age  uint
}

func main() {
	p := person{name: "farid", age: 26}
	fmt.Println(p)
}
