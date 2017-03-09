package main

import "fmt"

func main() {
	var emptySlice = []string{}
	var simpleSlice = []string{"a", "b", "c"}
	var makeSlice = make([]int, 5, 10)
	emptySlice[0] = "1"
	//	fmt.Println("emptySlice", emptySlice)
	//	emptySlice[0] = "0"
	//	fmt.Println("emptySlice", emptySlice)
	fmt.Println("makeSlice initialised", makeSlice)
	fmt.Println("length of makeSlice", len(makeSlice))
	fmt.Println("initital capacity of makeSlice", cap(makeSlice))

	makeSlice = append(makeSlice, 2, 3, 4)

	fused := append(emptySlice, simpleSlice...)

	fmt.Println("\nafter adding elements to makeSlice")
	fmt.Println(makeSlice)
	fmt.Println("length", len(makeSlice))
	fmt.Println("capacity", cap(makeSlice))

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := arr[1:5]
	slice2 := arr[:5]
	slice3 := arr[1:]
	slice4 := arr[:]

	fmt.Println("\nsliced arrays \n", slice1, slice2, slice3, slice4)
	fmt.Println("ignore", simpleSlice, fused)
}
