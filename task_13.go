package main

import "fmt"

func main() {
	x := 6
	y := 5

	x = x + y
	y = x - y
	x = x - y

	fmt.Print("x:", x, "\ny:", y)
}
