package main

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		result int
	}{
		{12, 8, 4},
		{18, 24, 6},
		{55, 89, 1},
		{0, 1, 1},
		{0, 0, 0},
		{37, 37, 37},
	}
	for _, test := range tests {
		t.Run(
			fmt.Sprintf("GCD(%d, %d)", test.x, test.y),
			func(t *testing.T) {
				got := GCD(test.x, test.y)
				if got != test.result {
					t.Errorf("GCD(%d, %d) = %d; want %d", test.x, test.y, got, test.result)
				}
			},
		)
	}
}
