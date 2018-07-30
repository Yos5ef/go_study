package main

import "fmt"

func len(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	// m := map[string]string {
	// 	"name": "zhouasdas",
	// 	"course": "golang",
	// }
	// m2 := make(map[string]int)

	// fmt.Println(m)
	// fmt.Println(m2)

	// fmt.Println("bianli")

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// courseName := m["course"]
	// fmt.Println(courseName)

	// qwe, ok := m["age"]
	// fmt.Println(qwe, ok)

	// delete(m, "name")
	// fmt.Println(m)

	fmt.Println(len("abcabcbb"))
}


//map的key 除了slice map function的内建类型都可以作为key

