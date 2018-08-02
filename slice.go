package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7,8}
	fmt.Println("arr[2:6]=", arr[2:6]);
	fmt.Println("arr[2:]=", arr[2:]);
	fmt.Println("arr[:6]=", arr[:6]);
	fmt.Println("arr[:]=", arr[:]);
	s1 := arr[2:6]
	s2 := arr[:6]
	updateSlice(s1)
	fmt.Println(s1, s2)
	fmt.Println(arr)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
} 
