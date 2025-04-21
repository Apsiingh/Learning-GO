package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	phoneNo int64
}

// 	FOR SINGLE STRUCT
// func main() {
// 1st way to declare struct value
// name := person{"Akhand" ,"Pratap"}

// 2nd way to declare struct value
// name := person{firstName: "Rahul", lastName: "Singh"}
// fmt.Print(name)

// 3rd way to declare struct value
// var name person
// name.firstName = "komal"
// name.lastName = "singh"
// fmt.Println(name)
// %+v is used to print all the struct key and their value
// fmt.Printf("%+v",name)

// }

//FOR NESTED STRUCT
func main() {
	// Don't , forget to give commas ,
	akhand := person{
		firstName: "Akhand",
		lastName:  "PratapSingh",
		contact: contactInfo{
			email:   "a.p.singh5508@gmail.com",
			phoneNo: 123456,
		},
	}
	// akhand.print()

	// &variable -- Give me the memory address of the value this
	// this variable is pointing at

	// akhandPointer := &akhand
	// akhandPointer.updateName("Rahul", "kumar")

	akhand.updateName("rahul", "singh")
	akhand.print()

	// OVER HERE WE NOT HAVE TO SPECIFY THE POINTER IN THE SLICE
	// BUT WHY WE HAVE TO SPECIFY POINTER  INSIDE THE STRUCT
	// watch video no :-- 012 Reference vs Value Types

	// Value Types -- Use pointers to change these things in a func
	//  int , float , string , bool , structs
	// Reference Types -- Don't worry about pointers with these
	// slices , maps , channels , pointers , functions

	goPlayGround()
}

func (pointerToPerson *person) updateName(firstName string, lastName string) {
	// *pointer -- Give me the value this memory address is pointing at
	(*pointerToPerson).firstName = firstName
	(*pointerToPerson).lastName = lastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func goPlayGround() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
	fmt.Println(s)
}
