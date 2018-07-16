package main

import "fmt"

func printArray(arr *[5]int) { //传长度为5的数组
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100;
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(&arr1)

	printArray(&arr3)
}

// var arr1 [5]int
// arr2 := [3]int{1, 3, 5}
// arr3 := [...]int{2, 4, 6, 8, 10}
// var grid [4][5]int   4代表四个数组 5代表每个数组长度为5
// range 第一个参数是下标 第二个参数是值
// [10]int 和 [20]int 是不同类型
// 调用func f(arr [10]int) 会拷贝数组
