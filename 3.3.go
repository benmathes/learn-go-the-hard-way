package main

import "fmt"

type Position struct {
	x int
	y int
}

func main() {
	pos := Position{1, 3}
	fmt.Printf("(%d,%d)\n", pos.x, pos.y)
}
