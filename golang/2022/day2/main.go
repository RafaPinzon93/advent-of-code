package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("Hello from the goroutine!")
}

func main() {
	go print()
	fmt.Println("Hello from the main function!")
	time.Sleep(1 * time.Second) // You need to specify the duration e.g., 1 * time.Second
}
