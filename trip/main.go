package main

import "fmt"

func trip(ar []bool, trip int) []bool {
	for i := 0; i < trip; i++ {
		for i := range ar {
			ar[i] = !ar[i]
		}
		for i := range ar {
			if (i+1)%2 == 0 {
				ar[i] = !ar[i]
			}
		}
		for i := range ar {
			if (i+1)%3 == 0 {
				ar[i] = !ar[i]
			}
		}
	}
	return ar
}

func lamp(sum int) []bool {
	ar := make([]bool, 100)
	for i := 0; i < 100; i++ {
		ar[i] = false
	}
	return ar
}

func main() {
	lamp := lamp(100)
	result := trip(lamp, 100)
	fmt.Println(result)
}
