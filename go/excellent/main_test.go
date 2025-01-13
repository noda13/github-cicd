package main

import (
	"fmt"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2, "Even"},
		{3, "Odd"},
		{0, "Even"},
		{-1, "Odd"},
		{-2, "Even"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%d", test.input), func(t *testing.T) {
			result := EvenOrOdd(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
