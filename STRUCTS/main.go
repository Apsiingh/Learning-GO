package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

type contactInfo struct {
	email string
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


//FOR INTERNAL STRUCT
func main() {
	// Don't , forget to give commas ,
	akhand := person{
		firstName:"Akhand",
		lastName:"PratapSingh",
		contact : contactInfo{
         email:"a.p.singh5508@gmail.com",
	     phoneNo:123456,
		},
	}
	// akhand.print()
 
	// &variable -- Give me the memory address of the value this 
	// this variable is pointing at
	akhandPointer := &akhand 
	akhandPointer.updateName("Rahul")
	akhand.print()
}


func (pointerToPerson *person) updateName (newName string){
    // *pointer -- Give me the value this memory address is pointing at 
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v",p)
}