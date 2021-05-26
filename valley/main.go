package main

import "fmt"

func valley(num int, steps []string) int {
	height := 0
	result := 0
	for _, e := range steps {
		if e == "U" {
			if height == -1 {
				result++
			}
			height++
		} else if e == "D" {
			height--
		}
	}
	return result
}

func main() {
	steps := []string{"U", "D", "D", "D", "U", "D", "U", "U"}
	valley := valley(len(steps), steps)
	fmt.Println(valley)
}
