package main

import (
	"testing"
)

func TestGetMaxJoltage(t *testing.T) {
	t.Run("Should return maximum joltage", func(t *testing.T) {
		test_cases := map[string]int{
			"987654321111111": 98,
			"811111111111119": 89,
			"234234234234278": 78,
			"818181911112111": 92,
		}

		for input, expected := range test_cases {
			result := GetMaximumJoltage(input)
			if result != expected {
				t.Errorf("expected %#v, but got %#v", expected, result)
			}
		}
	})
}

func TestGetSumMaxJoltage(t *testing.T) {
	t.Run("Should return maximum joltage sum", func(t *testing.T) {
		input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
		var expected int = 357
		sum := GetSumMaxJoltage(input)

		if sum != expected {
			t.Errorf("expected %d, but got %d", expected, sum)
		}
	},
	)
}
