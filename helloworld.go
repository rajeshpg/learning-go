package main

import (
	"fmt"

	"github.com/rajeshpg/learning-go/structs"
)

func main() {
	fmt.Println("hello gophers !!!")
	createPerson()
}

func createPerson() {
	p := structs.Person{Age: 36, Fname: "Rajesh", Lname: "Pitty"}
	p.Fullname()
	fmt.Println("full name ", p.FullName)

}

type Animal interface {
	Speak()
}

type Human interface {
	Speak()
}

type Humimal struct{}

func (h *Humimal) Speak() {
	fmt.Println("bow")
}
