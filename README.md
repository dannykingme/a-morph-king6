# Simple Calculator

A command-line calculator application written in Go, migrated from C++. This interactive menu-driven calculator performs basic arithmetic operations (addition, subtraction, multiplication, and division) with robust input validation and error handling.

## Overview

This project is a modernization of a C++ calculator application, preserving the original CLI interaction model while leveraging Go's simplicity, strong standard library, and idiomatic error handling. The migration maintains complete functional equivalence with the original C++ version, including identical menu options, prompts, result formatting, and error messages.

## Features

- **Interactive Menu System**: User-friendly menu with clear options (1-5)
- **Four Arithmetic Operations**: Addition, subtraction, multiplication, and division
- **Robust Input Validation**: Handles invalid menu choices and non-numeric input gracefully
- **Division-by-Zero Protection**: Prevents undefined behavior with clear error messaging
- **Floating-Point Support**: Works with both integers and decimal numbers using `float64`
- **Error Recovery**: Application continues running after errors, allowing users to retry
- **Formatted Output**: Results displayed in the format `num1 operator num2 = result`

## Requirements

- **Go 1.21+**: This project requires Go version 1.21 or higher
- **Go Modules**: Enabled by default in Go 1.16+

## Installation

1. Clone the repository or navigate to the project directory:
   ```bash
   cd /path/to/calculator
   ```

2. Download dependencies (if any):
   ```bash
   go mod download
   ```

## Building

Build the calculator executable:

```bash
go build -o calculator ./cmd/calculator
```

This creates an executable named `calculator` in the current directory.

## Running

Run the calculator:

```bash
./calculator
```

The application will display a welcome message and present an interactive menu:

```
Welcome to the Simple Calculator!

========== Simple Calculator ==========
1. Addition (+)
2. Subtraction (-)
3. Multiplication (*)
4. Division (/)
5. Exit
=======================================
Enter your choice (1-5):
```

## Usage Example

Here's a sample interaction session:

```
Welcome to the Simple Calculator!

========== Simple Calculator ==========
1. Addition (+)
2. Subtraction (-)
3. Multiplication (*)
4. Division (/)
5. Exit
=======================================
Enter your choice (1-5): 1
Enter first number: 2.5
Enter second number: 3.7
2.5 + 3.7 = 6.2

Press Enter to continue...

========== Simple Calculator ==========
1. Addition (+)
2. Subtraction (-)
3. Multiplication (*)
4. Division (/)
5. Exit
=======================================
Enter your choice (1-5): 4
Enter first number: 10
Enter second number: 0
Error: Division by zero is not allowed!

Press Enter to continue...

========== Simple Calculator ==========
1. Addition (+)
2. Subtraction (-)
3. Multiplication (*)
4. Division (/)
5. Exit
=======================================
Enter your choice (1-5): 5
Thank you for using the calculator! Goodbye!
```

## Project Structure

The project follows Go's standard project layout conventions:

```
.
├── cmd/
│   └── calculator/
│       ├── main.go           # Application entry point and main event loop
│       └── main_test.go      # Integration tests for complete CLI workflows
├── internal/
│   ├── calculator/
│   │   ├── core.go           # Pure arithmetic functions (Add, Subtract, Multiply, Divide)
│   │   └── core_test.go      # Unit tests for calculator operations
│   ├── errors/
│   │   ├── handler.go        # Error definitions and messages
│   │   └── handler_test.go   # Tests for error handling
│   ├── input/
│   │   ├── validation.go     # Input parsing and validation logic
│   │   └── validation_test.go # Tests for input validation
│   └── ui/
│       ├── menu.go           # Menu display and result formatting
│       └── menu_test.go      # Tests for UI components
├── go.mod                    # Go module configuration
├── go.sum                    # Dependency checksums
└── README.md                 # This file
```

### Package Descriptions

- **`cmd/calculator`**: Main application entry point. Contains the main event loop that orchestrates user interaction, menu display, input collection, operation dispatch, and result display.

- **`internal/calculator`**: Core calculator logic. Implements pure functions for the four arithmetic operations, including division-by-zero validation. All operations use `float64` to match C++ `double` semantics.

- **`internal/ui`**: User interface components. Handles menu rendering, welcome/goodbye messages, result formatting, and "Press Enter to continue" prompts. All UI text is defined as constants for easy maintenance.

- **`internal/input`**: Input validation and parsing. Provides functions to read and validate menu choices (1-5 range) and numeric operands, with clear error messages for invalid input.

- **`internal/errors`**: Error definitions. Centralizes all error types and messages used throughout the application, ensuring consistency with the original C++ error messages.

## Testing

The project includes comprehensive unit and integration tests to ensure correctness and preserve the CLI interaction contract from the C++ version.

### Run All Tests

```bash
go test ./...
```

### Run Tests with Verbose Output

```bash
go test ./... -v
```

### Run Tests for a Specific Package

```bash
# Test calculator operations
go test ./internal/calculator -v

# Test input validation
go test ./internal/input -v

# Test integration workflows
go test ./cmd/calculator -v
```

### Test Coverage

- **Unit Tests**: Validate individual components (arithmetic operations, input parsing, error handling, UI formatting)
- **Integration Tests**: Validate complete CLI workflows by simulating stdin/stdout interactions
  - Addition, subtraction, multiplication, division workflows
  - Division-by-zero error handling
  - Invalid menu choice and number input validation
  - Exit workflow and goodbye message
  - Multiple operations in sequence
  - Floating-point number handling

All tests verify that output messages match the C++ version to ensure functional equivalence.

## Design Decisions

### Error Handling

The Go implementation uses idiomatic error return values instead of C++ exceptions. The `Divide` function returns `(float64, error)`, and the main loop checks for errors and displays appropriate messages. This makes error handling explicit and follows Go best practices.

### Input Validation

Input reading uses `bufio.Scanner` for line-based reading, providing better control over input handling compared to C++'s stream-based approach. Invalid input is parsed with `strconv` functions, returning clear error messages that match the C++ version.

### Numeric Output Formatting

Results are formatted using `fmt.Printf` with the `%g` format verb, which approximates C++'s default double formatting by automatically choosing between fixed and exponential notation and removing trailing zeros.

### Package Organization

The project uses a single file per package structure, keeping the codebase simple and maintainable while following Go conventions. The `internal/` directory ensures implementation packages are not imported by external projects.

## Migration from C++

This Go implementation preserves the complete CLI interaction contract from the original C++ version:

- ✅ Same menu options (1-5) with equivalent labels
- ✅ Equivalent prompts for numeric input and choices
- ✅ Identical success and error messages
- ✅ Same control flow: loop until exit, input validation, and error recovery
- ✅ Numerically equivalent results with `float64` matching C++ `double` semantics
- ✅ Same user experience: welcome/goodbye messages, "Press Enter" pauses, operation flow

Minor differences from the C++ version:

- **Error Handling**: Uses Go's explicit error returns instead of exceptions
- **Input Reading**: Uses `bufio.Scanner` instead of C++ streams
- **Output Formatting**: Uses `%g` format verb instead of C++ stream default formatting (may produce minor formatting differences like omission of trailing zeros)

## Contributing

When modifying this project, please ensure:

1. All tests pass: `go test ./...`
2. New functionality includes appropriate unit and integration tests
3. Error messages remain consistent with the C++ version
4. CLI interaction contract is preserved
5. Code follows Go best practices and idioms

## License

This project is a migration exercise demonstrating C++ to Go modernization.

## Source

Original C++ implementation: https://github.com/dannykingme/a-morph-king6.git
