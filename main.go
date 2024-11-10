package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Now you have %g problems.\n", add(42, 13))
}
