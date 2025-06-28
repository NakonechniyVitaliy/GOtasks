package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person = Person{"Vitalii", 22}

	fmt.Println(p)
}
