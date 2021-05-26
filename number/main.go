package main

import (
	"fmt"
	"strconv"
	"strings"
)

func number(number int) string {
	str := strconv.Itoa(number)
	length := len(str)
	result := ""
	for i := 0; i < length; i++ {
		result += string(str[i]) + strings.Repeat("0", length-(i+1)) + "\n"
	}
	return result
}

func main() {
	num := number(1345679)
	fmt.Println(num)
}
