package main

import (
	"fmt"
	"math"
	"reflect"

	"rajeshpg.com/learning-go/interfaces"
	"rajeshpg.com/learning-go/structs"
)

func functions() {
	var a, b = SwapInts(1, 2)
	fmt.Println(a, b)
}

func SwapInts(x, y int) (i, j int) {
	i = y
	j = x
	return
}

func add(x, y int) int {
	return x + y
}

func makeFullName(person structs.Person) string {
	return person.Fname + " " + person.Lname
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) ID() float64 {
	return 1.1
}

func areaOfShape(s interfaces.Shape) {
	fmt.Println(reflect.TypeOf(s), s.Area())
}

// t, ok := s.(interfaces.Shape)
// fmt.Println(t, ok)

func main() {
	circle := Circle{10.0}
	areaOfShape(circle)
}
