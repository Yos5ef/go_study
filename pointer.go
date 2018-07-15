package main

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
} 

func swapTwo(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	c, d := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swapTwo(c, d))
}
