package structs

import "fmt"

type Person struct {
	Age      int
	Fname    string
	Lname    string
	salary   int
	FullName string
	// dept     Department

}

type Department struct {
	name string
}

func CreatePerson() {
	p1 := Person{36, "Rajesh", "Pitty", 500, ""}
	p2 := Person{Age: 36, Fname: "Rajesh"}

	p3 := new(Person)
	p3.Fname = "Rajesh"
	fmt.Println(p1, p2, p3)
}

func (person *Person) Fullname() string {
	person.FullName = person.Fname + " " + person.Lname
	return person.FullName
}
