package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	var err error
	delay := 1

	for i := 0; i < 3; i++ {

		err = doOperation()

		if err == nil {
			fmt.Println("Operation successful")
			return
		}

		fmt.Println("Error:", err)

		if i < 2 {
			fmt.Println("Retrying...")
			time.Sleep(time.Duration(delay) * time.Second)
			delay *= 2
		}
	}

	fmt.Println("Operation failed.")
}

func doOperation() error {
	return errors.New("loading again...")
}