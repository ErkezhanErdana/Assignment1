package main

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited: %.2f. New balance: %.2f\n", amount, b.Balance)
	} else {
		fmt.Println("Invalid deposit amount.")
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 {
		if b.Balance >= amount {
			b.Balance -= amount
			fmt.Printf("Withdrawn: %.2f. New balance: %.2f\n", amount, b.Balance)
		} else {
			fmt.Println("Insufficient balance.")
		}
	} else {
		fmt.Println("Invalid withdraw amount.")
	}
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, transaction := range transactions {
		if transaction > 0 {
			account.Deposit(transaction)
		} else if transaction < 0 {
			account.Withdraw(-transaction)
		}
	}
}

func main() {
	account := BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       0,
	}

	for {
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter deposit amount: ")
			var amount float64
			fmt.Scan(&amount)
			account.Deposit(amount)

		case 2:
			fmt.Print("Enter withdraw amount: ")
			var amount float64
			fmt.Scan(&amount)
			account.Withdraw(amount)

		case 3:
			account.GetBalance()

		case 4:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
