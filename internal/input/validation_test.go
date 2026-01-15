package input

import (
	"errors"
	"strings"
	"testing"
)

func TestReadMenuChoice(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantChoice  int
		wantErr     error
		errContains string
	}{
		{
			name:       "valid choice 1",
			input:      "1\n",
			wantChoice: 1,
			wantErr:    nil,
		},
		{
			name:       "valid choice 2",
			input:      "2\n",
			wantChoice: 2,
			wantErr:    nil,
		},
		{
			name:       "valid choice 3",
			input:      "3\n",
			wantChoice: 3,
			wantErr:    nil,
		},
		{
			name:       "valid choice 4",
			input:      "4\n",
			wantChoice: 4,
			wantErr:    nil,
		},
		{
			name:       "valid choice 5",
			input:      "5\n",
			wantChoice: 5,
			wantErr:    nil,
		},
		{
			name:       "valid choice with whitespace",
			input:      "  3  \n",
			wantChoice: 3,
			wantErr:    nil,
		},
		{
			name:       "choice too low (0)",
			input:      "0\n",
			wantChoice: 0,
			wantErr:    ErrInvalidChoice,
		},
		{
			name:       "choice too high (6)",
			input:      "6\n",
			wantChoice: 0,
			wantErr:    ErrInvalidChoice,
		},
		{
			name:       "negative choice",
			input:      "-1\n",
			wantChoice: 0,
			wantErr:    ErrInvalidChoice,
		},
		{
			name:       "very large number",
			input:      "999\n",
			wantChoice: 0,
			wantErr:    ErrInvalidChoice,
		},
		{
			name:       "non-numeric input",
			input:      "abc\n",
			wantChoice: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "empty input",
			input:      "\n",
			wantChoice: 0,
			wantErr:    ErrEmptyInput,
		},
		{
			name:       "whitespace only",
			input:      "   \n",
			wantChoice: 0,
			wantErr:    ErrEmptyInput,
		},
		{
			name:       "decimal number",
			input:      "2.5\n",
			wantChoice: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "input with letters",
			input:      "1a\n",
			wantChoice: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "special characters",
			input:      "!@#\n",
			wantChoice: 0,
			wantErr:    ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			gotChoice, gotErr := ReadMenuChoice(reader)

			if gotChoice != tt.wantChoice {
				t.Errorf("ReadMenuChoice() choice = %v, want %v", gotChoice, tt.wantChoice)
			}

			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("ReadMenuChoice() error = nil, want %v", tt.wantErr)
				} else if !errors.Is(gotErr, tt.wantErr) {
					t.Errorf("ReadMenuChoice() error = %v, want %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr != nil {
					t.Errorf("ReadMenuChoice() error = %v, want nil", gotErr)
				}
			}
		})
	}
}

func TestReadNumber(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantNumber float64
		wantErr    error
	}{
		{
			name:       "positive integer",
			input:      "42\n",
			wantNumber: 42.0,
			wantErr:    nil,
		},
		{
			name:       "negative integer",
			input:      "-15\n",
			wantNumber: -15.0,
			wantErr:    nil,
		},
		{
			name:       "positive decimal",
			input:      "3.14\n",
			wantNumber: 3.14,
			wantErr:    nil,
		},
		{
			name:       "negative decimal",
			input:      "-2.5\n",
			wantNumber: -2.5,
			wantErr:    nil,
		},
		{
			name:       "zero",
			input:      "0\n",
			wantNumber: 0.0,
			wantErr:    nil,
		},
		{
			name:       "zero as decimal",
			input:      "0.0\n",
			wantNumber: 0.0,
			wantErr:    nil,
		},
		{
			name:       "number with leading whitespace",
			input:      "  42  \n",
			wantNumber: 42.0,
			wantErr:    nil,
		},
		{
			name:       "scientific notation",
			input:      "1.5e2\n",
			wantNumber: 150.0,
			wantErr:    nil,
		},
		{
			name:       "very large number",
			input:      "999999999.999\n",
			wantNumber: 999999999.999,
			wantErr:    nil,
		},
		{
			name:       "very small decimal",
			input:      "0.00001\n",
			wantNumber: 0.00001,
			wantErr:    nil,
		},
		{
			name:       "non-numeric input",
			input:      "abc\n",
			wantNumber: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "empty input",
			input:      "\n",
			wantNumber: 0,
			wantErr:    ErrEmptyInput,
		},
		{
			name:       "whitespace only",
			input:      "   \n",
			wantNumber: 0,
			wantErr:    ErrEmptyInput,
		},
		{
			name:       "input with letters",
			input:      "12.5abc\n",
			wantNumber: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "multiple decimal points",
			input:      "1.2.3\n",
			wantNumber: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "special characters",
			input:      "!@#$\n",
			wantNumber: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "just a minus sign",
			input:      "-\n",
			wantNumber: 0,
			wantErr:    ErrInvalidInput,
		},
		{
			name:       "just a plus sign",
			input:      "+\n",
			wantNumber: 0,
			wantErr:    ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			gotNumber, gotErr := ReadNumber(reader)

			if gotNumber != tt.wantNumber {
				t.Errorf("ReadNumber() number = %v, want %v", gotNumber, tt.wantNumber)
			}

			if tt.wantErr != nil {
				if gotErr == nil {
					t.Errorf("ReadNumber() error = nil, want %v", tt.wantErr)
				} else if !errors.Is(gotErr, tt.wantErr) {
					t.Errorf("ReadNumber() error = %v, want %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr != nil {
					t.Errorf("ReadNumber() error = %v, want nil", gotErr)
				}
			}
		})
	}
}

// TestReadMenuChoiceEdgeCases tests additional edge cases
func TestReadMenuChoiceEdgeCases(t *testing.T) {
	// Test empty reader (no input at all)
	t.Run("no input available", func(t *testing.T) {
		reader := strings.NewReader("")
		_, err := ReadMenuChoice(reader)
		if err == nil {
			t.Error("ReadMenuChoice() with empty reader should return error")
		}
	})

	// Test boundary values
	t.Run("boundary value 1", func(t *testing.T) {
		reader := strings.NewReader("1\n")
		choice, err := ReadMenuChoice(reader)
		if err != nil || choice != 1 {
			t.Errorf("ReadMenuChoice(1) = %v, %v, want 1, nil", choice, err)
		}
	})

	t.Run("boundary value 5", func(t *testing.T) {
		reader := strings.NewReader("5\n")
		choice, err := ReadMenuChoice(reader)
		if err != nil || choice != 5 {
			t.Errorf("ReadMenuChoice(5) = %v, %v, want 5, nil", choice, err)
		}
	})
}

// TestReadNumberEdgeCases tests additional edge cases for number reading
func TestReadNumberEdgeCases(t *testing.T) {
	// Test empty reader (no input at all)
	t.Run("no input available", func(t *testing.T) {
		reader := strings.NewReader("")
		_, err := ReadNumber(reader)
		if err == nil {
			t.Error("ReadNumber() with empty reader should return error")
		}
	})

	// Test boundary values
	t.Run("maximum float64", func(t *testing.T) {
		reader := strings.NewReader("1.7976931348623157e308\n")
		num, err := ReadNumber(reader)
		if err != nil {
			t.Errorf("ReadNumber() with max float64 returned error: %v", err)
		}
		if num != 1.7976931348623157e308 {
			t.Errorf("ReadNumber() = %v, want max float64", num)
		}
	})

	t.Run("minimum positive float64", func(t *testing.T) {
		reader := strings.NewReader("2.2250738585072014e-308\n")
		num, err := ReadNumber(reader)
		if err != nil {
			t.Errorf("ReadNumber() with min positive float64 returned error: %v", err)
		}
		if num != 2.2250738585072014e-308 {
			t.Errorf("ReadNumber() = %v, want min positive float64", num)
		}
	})
}
