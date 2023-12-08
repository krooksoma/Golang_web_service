package main

import (
	"fmt"
)


func main(){
	var jim person = person{firstName: "Jim",
	 						lastName: "Thomas",
	 						contact: contactInfo{phone:" 234-222-1133", zipCode:"34234"}}
	// & gives the memory address to the variable it is pointing @
	jimPointer := &jim
	jimPointer.updatelastName("Jordan")
	jim.printInformation()

	//jim.updateLastName() would also work
}

func (p person) printInformation(){
	fmt.Printf("%+v", p)
}

// *pointer get this memory address and provides value in the address
// *<type> working with a pointer for that type
// *<variable> workinh to manipulate the value the pointer is referencing
func (pointerToPerson *person) updatelastName(name string){
	fmt.Print("inside update name")
	(*pointerToPerson).lastName = name
}