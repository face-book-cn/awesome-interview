package main

import "fmt"

func main() {
	i := hammingDistance(1,10)
	fmt.Println(i)
}

func hammingDistance(x int, y int) int {
	xor := x ^ y
	distance := 0
	for xor != 0 {
		if (xor & 1) == 1 {
			distance++
		}
		xor = xor >> 1
	}
	return distance
}