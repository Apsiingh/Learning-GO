package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 17}

	for i, n := range number {
		if n%2 == 0 {
			fmt.Println("Even number:", n, "at index:", i)
		} else {
			fmt.Println("Odd number:", n, "at index:", i)
		}
	}
}
