package main

import (
	"fmt"
)

func sock(ar []int, n int) int {
	calssified := make(map[int]int)
	for i := 0; i < n; i++ {
		calssified[ar[i]]++
	}
	result := 0
	for _, i := range calssified {
		result += i / 2
	}
	return result
}

func main() {
	ar := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := len(ar)
	fmt.Println(sock(ar, n))
}
