package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	revenue := askForInput("Enter revenue: ")
	expenses := askForInput("Enter expenses: ")
	taxRate := askForInput("Enter tax rate: ")

	ebt, profit, ratio := calculations(revenue, expenses, taxRate)

	printResults(ebt, profit, ratio)
	waitForEnter()
}

func askForInput(text string) float64 {
	var input float64
	fmt.Println(text)
	fmt.Scanln(&input)
	return input
}

func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
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
