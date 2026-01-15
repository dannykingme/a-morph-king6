package errors

import "testing"

// TestErrorMessages verifies that error messages match the C++ implementation exactly
func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "division by zero error",
			err:      ErrDivisionByZero,
			expected: "Division by zero is not allowed!",
		},
		{
			name:     "invalid input error",
			err:      ErrInvalidInput,
			expected: "Invalid input! Please enter a valid number.",
		},
		{
			name:     "invalid menu input error",
			err:      ErrInvalidMenuInput,
			expected: "Invalid input! Please enter a number between 1-5.",
		},
		{
			name:     "invalid choice error",
			err:      ErrInvalidChoice,
			expected: "Invalid choice! Please select a number between 1-5.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("Error message = %q, want %q", tt.err.Error(), tt.expected)
			}
		})
	}
}

// TestErrorsAreNotNil verifies that all error variables are properly initialized
func TestErrorsAreNotNil(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{"ErrDivisionByZero", ErrDivisionByZero},
		{"ErrInvalidInput", ErrInvalidInput},
		{"ErrInvalidMenuInput", ErrInvalidMenuInput},
		{"ErrInvalidChoice", ErrInvalidChoice},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Errorf("%s is nil, expected non-nil error", tt.name)
			}
		})
	}
}

// TestErrorsAreDistinct verifies that each error is a unique instance
func TestErrorsAreDistinct(t *testing.T) {
	errors := []error{
		ErrDivisionByZero,
		ErrInvalidInput,
		ErrInvalidMenuInput,
		ErrInvalidChoice,
	}

	// Check that no two errors are the same instance
	for i := 0; i < len(errors); i++ {
		for j := i + 1; j < len(errors); j++ {
			if errors[i] == errors[j] {
				t.Errorf("Errors at indices %d and %d are the same instance", i, j)
			}
		}
	}
}
