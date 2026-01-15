package calculator

import (
	"math"
	"testing"
)

// TestAdd tests the Add function with various inputs
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive numbers", 2.0, 3.0, 5.0},
		{"negative numbers", -5.0, -3.0, -8.0},
		{"mixed signs", 10.0, -3.0, 7.0},
		{"with zero", 5.0, 0.0, 5.0},
		{"zero with zero", 0.0, 0.0, 0.0},
		{"large numbers", 1e10, 2e10, 3e10},
		{"small numbers", 0.1, 0.2, 0.3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestSubtract tests the Subtract function with various inputs
func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive numbers", 10.0, 3.0, 7.0},
		{"negative numbers", -5.0, -3.0, -2.0},
		{"mixed signs", 10.0, -3.0, 13.0},
		{"with zero", 5.0, 0.0, 5.0},
		{"zero minus number", 0.0, 5.0, -5.0},
		{"same numbers", 5.0, 5.0, 0.0},
		{"large numbers", 1e10, 2e9, 8e9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestMultiply tests the Multiply function with various inputs
func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"positive numbers", 2.0, 3.0, 6.0},
		{"negative numbers", -5.0, -3.0, 15.0},
		{"mixed signs", 10.0, -3.0, -30.0},
		{"with zero", 5.0, 0.0, 0.0},
		{"zero with number", 0.0, 5.0, 0.0},
		{"with one", 5.0, 1.0, 5.0},
		{"large numbers", 1e5, 1e5, 1e10},
		{"small numbers", 0.1, 0.2, 0.02},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestDivide tests the Divide function with various inputs
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{"positive numbers", 10.0, 2.0, 5.0, false},
		{"negative numbers", -10.0, -2.0, 5.0, false},
		{"mixed signs", 10.0, -2.0, -5.0, false},
		{"with one", 5.0, 1.0, 5.0, false},
		{"zero divided by number", 0.0, 5.0, 0.0, false},
		{"large numbers", 1e10, 1e5, 1e5, false},
		{"small numbers", 0.2, 0.1, 2.0, false},
		{"division by zero", 10.0, 0.0, 0.0, true},
		{"zero by zero", 0.0, 0.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide(%v, %v) expected error, got nil", tt.a, tt.b)
				}
				if err != ErrDivisionByZero {
					t.Errorf("Divide(%v, %v) expected ErrDivisionByZero, got %v", tt.a, tt.b, err)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(got-tt.want) > 1e-9 {
					t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}

// TestDivideErrorMessage verifies the exact error message matches C++ version
func TestDivideErrorMessage(t *testing.T) {
	_, err := Divide(5.0, 0.0)
	if err == nil {
		t.Fatal("Expected error for division by zero, got nil")
	}
	expected := "Division by zero is not allowed!"
	if err.Error() != expected {
		t.Errorf("Error message = %q, want %q", err.Error(), expected)
	}
}
