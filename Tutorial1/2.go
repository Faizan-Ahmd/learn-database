package main

import (
	"fmt"
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}
func main() {
	si := []int{10, 30, 35, -11}
	fmt.Println(Index(si, 30))
	ss := []string{"Hello", "Faizan", "Found", "Ali"}
	fmt.Println(Index(ss, "Faizan"))
}