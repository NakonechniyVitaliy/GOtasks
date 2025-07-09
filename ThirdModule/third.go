package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

//  4
//	dataFromFile, err := os.ReadFile("people.json")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	personSlice := []Person{}
//	json.Unmarshal(dataFromFile, &personSlice)
//	fmt.Println(personSlice)

type Person struct {
	Name string
	Age  int
}

func main() {

	personFromFile, err := os.ReadFile("people.json")

	if err != nil {
		fmt.Println(err)
	}

	parsedPersons := []Person{}

	err = json.Unmarshal(personFromFile, &parsedPersons)

	if err != nil {
		fmt.Println(err)
	}

	for key, _ := range parsedPersons {
		parsedPersons[key].Age++
	}

	err = os.WriteFile("people.json", string(parsedPersons), 0644)

}
