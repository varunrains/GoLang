package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// custom type inside a custom side
//Embedding structs
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//varun := person{"Varun", "Upadhya"}
	//This is the better way of initializing with the name
	// varun := person{firstName: "Varun", lastName: "Upadhya"}
	// fmt.Println(varun)
	// fmt.Println("----------------------------")
	//thirdWayOfAccessingStructValues()

	varun := person{
		firstName: "Varun",
		lastName:  "Upadhya",
		contact:   contactInfo{email: "varun.upadhya1@gmail.com", zipCode: 560093},
	}
	//This will not work
	//Pointers in GO is straight forward
	// varun.updateName("New firstName")
	// varun.print()

	//To make it work we need to use pointers and ampersand

	// varunPointer := &varun
	// varunPointer.updateName("NewfirstName")
	// varun.print()

	//You can use pointer short cut to reduce the above lines of code
	varun.updateName("NewName")
	varun.print()

	//Big difference between structs and slices
	//When you pass struct a different method and then update it, it will be just pass by value and
	//the orginal struct will not be updated
	//But in slices the original value that was passed will be updated
	//One of the gotchas with pointers.

	slices := []string{"Hi", "There"}
	changeSliceValues(slices)
	fmt.Println(slices)
}

func changeSliceValues(sliceValue []string) {
	//Slice is a reference type so you don't have to
	//worry about the pointers
	//whereas struct is a value type
	sliceValue[1] = "World"

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func structWithReceiverFunctions() {

}

func thirdWayOfAccessingStructValues() {
	var varun person

	varun.firstName = "Varun"
	varun.lastName = "Upadhya"
	fmt.Println(varun)
	fmt.Printf("%+v", varun)
}
