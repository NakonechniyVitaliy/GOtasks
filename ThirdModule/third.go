package main

import "fmt"

func perimetrRectangle(firstSide, secondSide float32) (square float32) {
	square = firstSide * secondSide
	return
}

func printSquare(square float32) {
	fmt.Println(square)
}

func main() {

	a := float32(10)
	b := float32(20)

	printSquare(perimetrRectangle(a, b))

}
