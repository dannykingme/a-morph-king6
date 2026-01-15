package input

import (
	"bufio"
	"calculator/internal/errors"
	"io"
	"strconv"
	"strings"
)

// ReadMenuChoice reads a line from the provided reader, parses it as an integer,
// and validates that it's in the range 1-5.
// Returns the choice and nil on success, or 0 and an error on failure.
func ReadMenuChoice(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, errors.ErrInvalidMenuInput
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		return 0, errors.ErrInvalidMenuInput
	}

	choice, err := strconv.Atoi(line)
	if err != nil {
		return 0, errors.ErrInvalidMenuInput
	}

	if choice < 1 || choice > 5 {
		return 0, errors.ErrInvalidChoice
	}

	return choice, nil
}

// ReadNumber reads a line from the provided reader and parses it as a float64.
// Returns the number and nil on success, or 0 and an error on failure.
func ReadNumber(reader io.Reader) (float64, error) {
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, errors.ErrInvalidInput
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		return 0, errors.ErrInvalidInput
	}

	number, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 0, errors.ErrInvalidInput
	}

	return number, nil
}
