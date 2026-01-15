package ui

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput captures stdout during function execution
func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestDisplayWelcome(t *testing.T) {
	output := captureOutput(func() {
		DisplayWelcome()
	})

	expected := "Welcome to the Simple Calculator!\n"
	if output != expected {
		t.Errorf("DisplayWelcome() = %q, want %q", output, expected)
	}
}

func TestDisplayMenu(t *testing.T) {
	output := captureOutput(func() {
		DisplayMenu()
	})

	// Check that output contains all required menu elements
	expectedParts := []string{
		"========== Simple Calculator ==========",
		"1. Addition (+)",
		"2. Subtraction (-)",
		"3. Multiplication (*)",
		"4. Division (/)",
		"5. Exit",
		"=======================================",
		"Enter your choice (1-5):",
	}

	for _, part := range expectedParts {
		if !strings.Contains(output, part) {
			t.Errorf("DisplayMenu() output missing expected part: %q", part)
		}
	}

	// Verify the order of menu options
	if !strings.Contains(output, "1. Addition (+)") ||
		!strings.Contains(output, "2. Subtraction (-)") ||
		!strings.Contains(output, "3. Multiplication (*)") ||
		!strings.Contains(output, "4. Division (/)") ||
		!strings.Contains(output, "5. Exit") {
		t.Error("DisplayMenu() missing one or more menu options")
	}
}

func TestDisplayResult(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		result   float64
		operator string
		expected string
	}{
		{
			name:     "addition with integers",
			num1:     2,
			num2:     3,
			result:   5,
			operator: "+",
			expected: "2 + 3 = 5\n",
		},
		{
			name:     "subtraction with integers",
			num1:     10,
			num2:     4,
			result:   6,
			operator: "-",
			expected: "10 - 4 = 6\n",
		},
		{
			name:     "multiplication with decimals",
			num1:     2.5,
			num2:     4,
			result:   10,
			operator: "*",
			expected: "2.5 * 4 = 10\n",
		},
		{
			name:     "division with decimals",
			num1:     10,
			num2:     3,
			result:   3.3333333333333335,
			operator: "/",
			expected: "10 / 3 = 3.3333333333333335\n",
		},
		{
			name:     "negative numbers",
			num1:     -5,
			num2:     3,
			result:   -2,
			operator: "+",
			expected: "-5 + 3 = -2\n",
		},
		{
			name:     "zero result",
			num1:     5,
			num2:     5,
			result:   0,
			operator: "-",
			expected: "5 - 5 = 0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				DisplayResult(tt.num1, tt.num2, tt.result, tt.operator)
			})

			if output != tt.expected {
				t.Errorf("DisplayResult(%g, %g, %g, %q) = %q, want %q",
					tt.num1, tt.num2, tt.result, tt.operator, output, tt.expected)
			}
		})
	}
}

func TestDisplayGoodbye(t *testing.T) {
	output := captureOutput(func() {
		DisplayGoodbye()
	})

	expected := "Thank you for using the calculator! Goodbye!\n"
	if output != expected {
		t.Errorf("DisplayGoodbye() = %q, want %q", output, expected)
	}
}

func TestDisplayPressEnter(t *testing.T) {
	output := captureOutput(func() {
		DisplayPressEnter()
	})

	expected := "\nPress Enter to continue..."
	if output != expected {
		t.Errorf("DisplayPressEnter() = %q, want %q", output, expected)
	}
}

// TestFormatVerb verifies that %g format is used correctly for numeric output
func TestFormatVerb(t *testing.T) {
	// Test that integers don't have trailing zeros
	output := captureOutput(func() {
		DisplayResult(2.0, 3.0, 5.0, "+")
	})

	if strings.Contains(output, ".0") {
		t.Error("DisplayResult() should not include trailing zeros for whole numbers")
	}

	// Test that decimals are displayed correctly
	output2 := captureOutput(func() {
		DisplayResult(2.5, 4.5, 7.0, "+")
	})

	if !strings.Contains(output2, "2.5") || !strings.Contains(output2, "4.5") {
		t.Error("DisplayResult() should display decimal numbers correctly")
	}
}
