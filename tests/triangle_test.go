package triangle

import (
	"testing"
	"triangle/internal"
)

func TestClassifyTriangle(t *testing.T) {
	tests := []struct {
		name     string
		a, b, c  int
		expected string
	}{
		// Invalid Input
		{
			name:     "Lower boundary invalid",
			a:        0,
			b:        5,
			c:        5,
			expected: "Invalid Input",
		},
		{
			name:     "Upper boundary invalid",
			a:        101,
			b:        5,
			c:        5,
			expected: "Invalid Input",
		},

		// Not a Triangle
		{
			name:     "a + b = c boundary",
			a:        1,
			b:        2,
			c:        3,
			expected: "Not a Triangle",
		},
		{
			name:     "a + c < b",
			a:        1,
			b:        5,
			c:        1,
			expected: "Not a Triangle",
		},
		{
			name:     "b + c < a",
			a:        5,
			b:        1,
			c:        1,
			expected: "Not a Triangle",
		},

		// Equilateral
		{
			name:     "Lower valid boundary equilateral",
			a:        1,
			b:        1,
			c:        1,
			expected: "Equilateral",
		},
		{
			name:     "Upper valid boundary equilateral",
			a:        100,
			b:        100,
			c:        100,
			expected: "Equilateral",
		},

		// Isosceles
		{
			name:     "a = b",
			a:        5,
			b:        5,
			c:        3,
			expected: "Isosceles",
		},
		{
			name:     "a = c",
			a:        5,
			b:        3,
			c:        5,
			expected: "Isosceles",
		},
		{
			name:     "b = c",
			a:        3,
			b:        5,
			c:        5,
			expected: "Isosceles",
		},

		// Scalene
		{
			name:     "Scalene triangle",
			a:        4,
			b:        5,
			c:        6,
			expected: "Scalene",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := internal.ClassifyTriangle(tc.a, tc.b, tc.c)

			if result != tc.expected {
				t.Errorf(
					"ClassifyTriangle(%d, %d, %d) = %s; want %s",
					tc.a,
					tc.b,
					tc.c,
					result,
					tc.expected,
				)
			}
		})
	}
}