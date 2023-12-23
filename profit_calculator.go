package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Println("Enter revenue: ")
	fmt.Scanln(&revenue)

	fmt.Println("Enter expenses: ")
	fmt.Scanln(&expenses)

	fmt.Println("Enter tax rate: ")
	fmt.Scanln(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Your Earnings Before Tax (EBT): ")
	fmt.Println(ebt)

	fmt.Println("Your Profit: ")
	fmt.Println(profit)

	fmt.Println("Your Ratio: ")
	fmt.Println(ratio)

	fmt.Println("Press Enter to exit...")
	waitForEnter()
}

func waitForEnter() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Wait for input
}
