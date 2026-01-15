package main

import (
	"bufio"
	"calculator/internal/calculator"
	"calculator/internal/errors"
	"calculator/internal/input"
	"calculator/internal/ui"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a single scanner for stdin to avoid buffering issues
	scanner := bufio.NewScanner(os.Stdin)

	// Display welcome message
	ui.DisplayWelcome()

	// Main event loop
	for {
		// Display menu
		ui.DisplayMenu()

		// Read and validate menu choice
		choice, err := readMenuChoice(scanner)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Check for exit condition
		if choice == 5 {
			ui.DisplayGoodbye()
			break
		}

		// Read first number
		fmt.Print("Enter first number: ")
		num1, err := readNumber(scanner)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Read second number
		fmt.Print("Enter second number: ")
		num2, err := readNumber(scanner)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Perform operation based on choice and display result
		var result float64
		var opErr error

		switch choice {
		case 1:
			result = calculator.Add(num1, num2)
			ui.DisplayResult(num1, num2, result, "+")
		case 2:
			result = calculator.Subtract(num1, num2)
			ui.DisplayResult(num1, num2, result, "-")
		case 3:
			result = calculator.Multiply(num1, num2)
			ui.DisplayResult(num1, num2, result, "*")
		case 4:
			result, opErr = calculator.Divide(num1, num2)
			if opErr != nil {
				// Handle division by zero error
				if opErr == errors.ErrDivisionByZero {
					fmt.Println("Error:", opErr.Error())
				} else {
					fmt.Println("Error:", opErr.Error())
				}
			} else {
				ui.DisplayResult(num1, num2, result, "/")
			}
		}

		// Display "Press Enter to continue" prompt and wait for Enter
		ui.DisplayPressEnter()
		scanner.Scan()
	}
}

// readMenuChoice reads and validates a menu choice using the provided scanner
func readMenuChoice(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, errors.ErrInvalidMenuInput
	}

	line := strings.TrimSpace(scanner.Text())
	return input.ReadMenuChoice(strings.NewReader(line + "\n"))
}

// readNumber reads and validates a number using the provided scanner
func readNumber(scanner *bufio.Scanner) (float64, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, errors.ErrInvalidInput
	}

	line := strings.TrimSpace(scanner.Text())
	return input.ReadNumber(strings.NewReader(line + "\n"))
}
