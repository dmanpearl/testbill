package main

// From a literal slice of 8 integers print the sum of each group of 3 values.

import (
	"fmt"
	"main/proc"
)

func main() {
	fmt.Println(proc.Process([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
}
