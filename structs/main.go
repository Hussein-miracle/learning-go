package main

import (
	"fmt"
	"reflect"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// // alex := person{firstName: "Alex", lastName: "Handerson"}
	// // fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Handerson"

	// fmt.Printf("%v", alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Wood",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1233,
		},
	}

	jim.print()
	// // jimPointer := (&jim)
	// // fmt.Println(&jimPointer, "jimAddress")
	// // jimPointer.updateName("Jam")
	// jim.updateName("Jam")
	// jim.print()

	fmt.Printf("                           ")
	values := reflect.ValueOf(jim)
	typeS := values.Type()
	types := reflect.TypeOf(jim)

	fmt.Println(values)
	fmt.Println(types)
	fmt.Println(typeS)
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(i)
		fmt.Println(types.Field(i).Index)
		fmt.Println(string(types.Field(i).Name), values.Field(i))

	}
	// mySlice := []string{"toy", "jug", "car", "card", "dream"}
	// fmt.Print(mySlice)
	// updateSlice(mySlice)
	// fmt.Print(mySlice)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateSlice(s []string) {
	s[0] = "gun"
}
