package genmath_test

import (
	"testing"

	"github.com/krelinga/go-genmath"
)

func TestMax(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		tests := []struct {
			a, b     int
			expected int
		}{
			{1, 2, 2},
			{5, 3, 5},
			{-1, -2, -1},
			{0, 0, 0},
		}
		for _, tt := range tests {
			result := genmath.Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		}
	})

	t.Run("Float64", func(t *testing.T) {
		tests := []struct {
			a, b     float64
			expected float64
		}{
			{1.5, 2.5, 2.5},
			{-1.1, -2.2, -1.1},
			{0.0, 0.0, 0.0},
		}
		for _, tt := range tests {
			result := genmath.Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		}
	})

	t.Run("String", func(t *testing.T) {
		tests := []struct {
			a, b     string
			expected string
		}{
			{"apple", "banana", "banana"},
			{"grape", "apple", "grape"},
			{"", "", ""},
			{"zebra", "ant", "zebra"},
		}
		for _, tt := range tests {
			result := genmath.Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%q, %q) = %q; want %q", tt.a, tt.b, result, tt.expected)
			}
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		tests := []struct {
			a, b     int
			expected int
		}{
			{1, 2, 1},
			{5, 3, 3},
			{-1, -2, -2},
			{0, 0, 0},
		}
		for _, tt := range tests {
			result := genmath.Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Min(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		}
	})

	t.Run("Float64", func(t *testing.T) {
		tests := []struct {
			a, b     float64
			expected float64
		}{
			{1.5, 2.5, 1.5},
			{-1.1, -2.2, -2.2},
			{0.0, 0.0, 0.0},
		}
		for _, tt := range tests {
			result := genmath.Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Min(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		}
	})

	t.Run("String", func(t *testing.T) {
		tests := []struct {
			a, b     string
			expected string
		}{
			{"apple", "banana", "apple"},
			{"grape", "apple", "apple"},
			{"", "", ""},
			{"zebra", "ant", "ant"},
		}
		for _, tt := range tests {
			result := genmath.Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Min(%q, %q) = %q; want %q", tt.a, tt.b, result, tt.expected)
			}
		}
	})
}
