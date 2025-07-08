package main

import (
	"fmt"
	"os"
)

// 	1
// 	data := []byte("Hello, File!")
// 	e := os.WriteFile("output.txt", data, 0777)
//
// 	if e != nil {
//   	fmt.Println("Error create/write file!\n", e)
// 	}

//  2
//	file_data, e := os.ReadFile("output.txt")
//
//	if e != nil {
//  	fmt.Println(e)
//	}
//
//	fmt.Println(string(file_data))

//  3

//
//
//	jsonSlice, e := json.Marshal(personSlice)
//
//	if e != nil {
//		fmt.Println("error:", e)
//	}
//
//	e = os.WriteFile("people.json", jsonSlice, 0644)
//
//	if e != nil {
//		fmt.Println("error:", e)
//	}

type Person struct {
	Name string
	Age  int
}

func main() {

	dataFromFile, err := os.ReadFile("people.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dataFromFile)

}
