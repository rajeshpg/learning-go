package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()
	//anonymous
	go numbers()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func numbers() {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	for number := 1; number < 27; number++ {
		fmt.Printf("%d ", number)
	}
}
