package main

import "fmt"

func main() {
	fmt.Println("Count down")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
