package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
)

func main() {

	revenue, revenueErr := askForInput("Enter revenue: ")

	if revenueErr != nil {
		fmt.Println("Please enter a valid revenue amount.")
		return
	}

	expenses, expensesErr := askForInput("Enter expenses: ")

	if expensesErr != nil {
		fmt.Println("Please enter a valid expenses amount.")
		return
	}

	taxRate, taxErr := askForInput("Enter tax rate: ")

	if taxErr != nil {
		fmt.Println("Please enter a valid tax rate.")
		return
	}

	ebt, profit, ratio := calculations(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)
	printResults(ebt, profit, ratio)
	waitForEnter()
}

func askForInput(text string) (float64, error) {
	var input float64
	fmt.Println(text)
	fmt.Scanln(&input)
	if input <= 0 {
		fmt.Println("Exiting program...")
		return 0, errors.New("invalid input")
	}
	return input, nil
}

func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("Your Earnings Before Tax (EBT): %.2f\nYour Profit: %.2f\nYour Ratio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func printResults(ebt, profit, ratio float64) {
	fmt.Printf("Your Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Your Profit: %.2f\n", profit)
	fmt.Printf("Your Ratio: %.2f\n", ratio)
}

func waitForEnter() {
	fmt.Println("Press enter to exit...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
