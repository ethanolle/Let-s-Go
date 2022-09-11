package main

import "fmt"

type contactInfo struct {
	address string
	number int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	ethan := person{
		firstName: "lol",
		lastName: "lol",
		contact: contactInfo{
			address: "lol",
			number: 123,
		},
	}
	ethan.print()
	ethan.updateLastName("sayagh")
	ethan.print()
}


func (ethanPointer *person) updateLastName(newLastName string) {
	(*ethanPointer).lastName = newLastName
}
func (p person) print(){
	fmt.Printf("%+v", p)
}
