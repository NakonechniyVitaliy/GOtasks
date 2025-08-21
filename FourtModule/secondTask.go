package FourtModule

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Sound string
}

type Cat struct {
	Sound string
}

func (d Dog) Speak() string {
	return d.Sound
}

func (c Cat) Speak() string {
	return c.Sound
}

func Speaking(animals []Animal) {
	for _, animal := range animals {
		fmt.Printf("\nsays %s", animal.Speak())
	}
}

func Second() {
	d := Dog{"Woof"}
	c := Cat{"Meow"}

	var sliceAnimals = []Animal{d, c}
	Speaking(sliceAnimals)

}
