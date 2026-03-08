package main

import (
	"errors"
	"fmt"
)

type bankAccount struct {
	Owner   string
	Balance float64
}

func main() {
	var name string
	var balance float64

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter starting balance: ")
	fmt.Scan(&balance)

	account := bankAccount{name, balance}

	for {
		fmt.Println("Choose operation")
		fmt.Println("1.Deposit")
		fmt.Println("2.Withdraw")
		fmt.Println("3.Show the balance")
		fmt.Println("4.Exit")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			var amount float64
			fmt.Println("Enter deposit amount:")
			fmt.Scan(&amount)
			err := account.Deposit(amount)
			if err != nil {
				fmt.Println("Error.", err)
			} else {
				fmt.Println("Deposit successful.")
			}
		case 2:
			var amount float64
			fmt.Println("Enter withdraw amount:")
			fmt.Scan(&amount)
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println("Error.", err)
			} else {
				fmt.Println("Withdraw successful.")
			}
		case 3:
			fmt.Println("Current balance: ", account.Balance)
		case 4:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Wrong choice")
		}
	}
}

var invalidAmount = errors.New("Amount must be greater than 0.")
var invalidwithdraw = errors.New("Not enough money.")

func (b *bankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return invalidAmount
	}
	b.Balance += amount
	return nil
}
func (b *bankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return invalidAmount
	}
	if amount > b.Balance {
		return invalidwithdraw
	}
	b.Balance -= amount
	return nil
}
