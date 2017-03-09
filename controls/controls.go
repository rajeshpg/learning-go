package main

import (
	"fmt"
	"os"
)

func simpleFor() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	for {
		fmt.Println("for'ever")
	}
}

func simpleIf(i int) {
	if i%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func ifWithStatement() {
	if f, err := os.Open("filePath"); err != nil {
		f.Close()
	}
}

func evenOrOdd(i int) {
	switch i%2 == 0 {
	case true:
		fmt.Println("even")
	case false:
		fmt.Println("odd")
	}
}
