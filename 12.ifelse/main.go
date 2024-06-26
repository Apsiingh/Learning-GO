package main

import "fmt"

func main() {
	fmt.Println("Learnig If and else in goLang")

	login := 25
	var result string
	if login > 25 {
		result = "more"
	} else if login < 25 {
		result = "less"
	} else {
		result = "same"
	}
	fmt.Println(result)

	num := 3
	if num < 5 {
		fmt.Println("less than")
	} else if num > 5 {
		fmt.Println("more than")
	} else {
		fmt.Println("save")
	}
	
	// The number is Even and Odd

	if number := 2; number%2 == 0 {
		fmt.Println("The Number is Even")
	} else {
		fmt.Println("The Number id Odd")
	}

}
