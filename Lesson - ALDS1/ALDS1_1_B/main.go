package main

import "fmt"

// 最大公約数は
func GCD(x, y int) int {
	if x > y {
		x, y = y, x
	}
	for y > 0 {
		r := x % y
		x = y
		y = r
	}
	return x
}

func main() {
	var x int
	var y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%d\n", GCD(x, y))
}
