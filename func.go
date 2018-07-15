package main

import (
	"fmt"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" + op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

// func apply(op func(int, int) int, a, b int) int {
// 	fmt.Printf("call")
// 	return op(a, b)
// }

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(3, 4, "*"))	
	a, b := div(13, 3)
	fmt.Println(a, b)
	fmt.Println(sum(1, 2, 3, 4, 5))
}


//函数可以返回多个值
