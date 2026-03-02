package main

import "fmt"

func main() {
	for {
		var input int
		fmt.Println("Input:")
		fmt.Scan(&input)
		if input%3 == 0 && input%5 == 0 {
			fmt.Println("FizzBuzz")
			break
		}
		if input%3 == 0 && input%5 != 0 {
			fmt.Println("Fizz")
			break
		}
		if input%3 != 0 && input%5 == 0 {
			fmt.Println("Buzz")
			break
		}
		if input%3 != 0 && input%5 != 0 {
			fmt.Println(input)
			break
		}
	}
}
