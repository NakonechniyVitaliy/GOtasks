package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//  2
//  func PrintPerson(p Person) {
//	  fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
//  }

//  3
//	func Birthday(p *Person) {
//	  p.Age++
//	}
//  Birthday(&p)

//  4
//  type Person struct {
//	  Name    string
//	  Age     int
//	  Address Address
//  }
//  type Address struct {
//    City, Street, Zip string
//  }
//
//  var p Person = Person{"Vitalii", 22, Address{"Kyiv", "Shepeleva", "09823"}}

//  5
//  func NewPerson(name string, age int) Person {
//	  return Person{name, age}
//  }

// 6 - 7
//  func (p Person) IsAdult() bool {
//	  return p.Age > 25
//  }
//  personSlice := []Person{
//	  {"Bob", 15},
//	  {"Alice", 18},
//	  {"Ivan", 26},
//  }
//
//  var adultPerson []Person
//
//  for _, person := range personSlice {
//	  if person.IsAdult() {
//		  adultPerson = append(adultPerson, person)
//	  }
//  }

//  8
//  func IsSamePerson(p1, p2 Person) bool {
//	  return p1.Name == p2.Name && p1.Age == p2.Age
//  }

//  9
//  personMap := map[string]Person{
//		"Bob":     Person{"Bob", 20},
//		"Vitalii": Person{"Vitalii", 22},
//		"Anya":    Person{"Anya", 24},
//  }
//
//  fmt.Println(personMap["Bob"])

//  10
//  func (p Person) String() (stringPerson string) {
//	  return fmt.Sprintf("%s (%d y.o.)", p.Name, p.Age)
//  }

func main() {
	p := Person{"Vitalii", 22}
	fmt.Println(p)
}
