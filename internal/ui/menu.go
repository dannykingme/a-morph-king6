package ui

import "fmt"

// UI text constants
const (
	welcomeMessage       = "Welcome to the Simple Calculator!"
	goodbyeMessage       = "Thank you for using the calculator! Goodbye!"
	menuSeparatorTop     = "\n========== Simple Calculator =========="
	menuSeparatorBottom  = "======================================="
	menuOption1          = "1. Addition (+)"
	menuOption2          = "2. Subtraction (-)"
	menuOption3          = "3. Multiplication (*)"
	menuOption4          = "4. Division (/)"
	menuOption5          = "5. Exit"
	menuPrompt           = "Enter your choice (1-5): "
	pressEnterPrompt     = "\nPress Enter to continue..."
)

// DisplayWelcome prints the welcome message to the user
func DisplayWelcome() {
	fmt.Println(welcomeMessage)
}

// DisplayMenu prints the calculator menu with options 1-5
func DisplayMenu() {
	fmt.Println(menuSeparatorTop)
	fmt.Println(menuOption1)
	fmt.Println(menuOption2)
	fmt.Println(menuOption3)
	fmt.Println(menuOption4)
	fmt.Println(menuOption5)
	fmt.Println(menuSeparatorBottom)
	fmt.Print(menuPrompt)
}

// DisplayResult formats and prints the calculation result
// Uses %g format verb to match C++ double output formatting
func DisplayResult(num1, num2, result float64, operator string) {
	fmt.Printf("%g %s %g = %g\n", num1, operator, num2, result)
}

// DisplayGoodbye prints the goodbye message to the user
func DisplayGoodbye() {
	fmt.Println(goodbyeMessage)
}

// DisplayPressEnter prints the "Press Enter to continue" prompt
func DisplayPressEnter() {
	fmt.Print(pressEnterPrompt)
}
