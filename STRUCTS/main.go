package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// name := person{firstName: "Komal", lastName: "Singh"}
	// fmt.Print(name)

	var name person
	fmt.Println(name)

}
