package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello, Channel!"
	}()

	fmt.Println(<-ch)
}
