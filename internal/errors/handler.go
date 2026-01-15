package errors

import "errors"

// Error definitions for the calculator application.
// These errors match the error messages from the C++ implementation.

var (
	// ErrDivisionByZero is returned when attempting to divide by zero.
	// Matches C++ message: "Division by zero is not allowed!" (calculator.cpp:22)
	ErrDivisionByZero = errors.New("Division by zero is not allowed!")

	// ErrInvalidInput is returned when the user enters non-numeric input.
	// Matches C++ message: "Invalid input! Please enter a valid number." (calculator.cpp:47, 54)
	ErrInvalidInput = errors.New("Invalid input! Please enter a valid number.")

	// ErrInvalidMenuInput is returned when the user enters non-numeric input for menu choice.
	// Matches C++ message: "Invalid input! Please enter a number between 1-5." (calculator.cpp:73)
	ErrInvalidMenuInput = errors.New("Invalid input! Please enter a number between 1-5.")

	// ErrInvalidChoice is returned when the user enters a menu choice outside the 1-5 range.
	// Matches C++ message: "Invalid choice! Please select a number between 1-5." (calculator.cpp:84)
	ErrInvalidChoice = errors.New("Invalid choice! Please select a number between 1-5.")
)
