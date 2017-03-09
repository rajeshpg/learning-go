package main

import "fmt"

func main() {
	simpleArray()
}

func simpleArray() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	var slice = arr[0:2]
	fmt.Println(slice)
}
