package main

import "fmt"

type Sport struct {
	name string
	positon string
}


type person struct {
	name string
	age  uint
	favsport []Sport
}

func main() {
	p1 := person{ age: 38, name: "Messi", favsport: []Sport{{"Football","Striker"}}}
	fmt.Println(p1.name, p1.age)
	fmt.Println(p1.favsport[0].name, p1.favsport[0].positon)
}
