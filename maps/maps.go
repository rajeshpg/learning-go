package main

import "fmt"

func main() {
	var intString map[int]string
	intString = make(map[int]string)

	intString[1] = "one"
	intString[2] = "two"
	intString[2] = "three"
	intString[2] = "four"

	for key, value := range intString {
		fmt.Printf("key: %d; value: %s", key, value)
	}

	if value, exists := intString[1]; exists {
		fmt.Println(value)
	}

}
