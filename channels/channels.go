package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("printer ", msg)
		time.Sleep(time.Second * 1)
	}
}

func printer2(c chan string) {
	for {
		msg := <-c
		fmt.Println("printer2 ", msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	go printer2(c)
	fromc := <-c
	fmt.Println("in main", fromc)
	var input string
	fmt.Scanln(&input)
}
