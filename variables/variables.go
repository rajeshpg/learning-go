package main

import "fmt"

func variables() {
	var x string = "hello" // HL
	var y = " world"       // HL
	println(x + y)

	var defaultString string // HL
	fmt.Println("default value for string is " + defaultString)

	z := "hello world" // HL
	fmt.Println(z)

	const v string = "i am immutable." // HL
	fmt.Println(v + " check it out")
	fmt.Println(v)
}
