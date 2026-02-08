package main

import (
	"testing"
)

func TestGetMaxJoltage(t *testing.T) {
	t.Run("Should return maximum joltage, 2 digits", func(t *testing.T) {
		test_cases := map[string]int{
			"987654321111111": 98,
			"811111111111119": 89,
			"234234234234278": 78,
			"818181911112111": 92,
		}

		for input, expected := range test_cases {
			result := GetMaximumJoltage(input, 2)
			if result != expected {
				t.Errorf("expected %#v, but got %#v", expected, result)
			}
		}
	})

	t.Run("Should return maximum joltage, 12 digits", func(t *testing.T) {
		test_cases := map[string]int{
			"987654321111111": 987654321111,
			"811111111111119": 811111111119,
			"234234234234278": 434234234278,
			"818181911112111": 888911112111,
		}

		for input, expected := range test_cases {
			result := GetMaximumJoltage(input, 12)
			if result != expected {
				t.Errorf("expected %#v, but got %#v", expected, result)
			}
		}
	})
}

func TestGetSumMaxJoltage(t *testing.T) {
	t.Run("Should return maximum joltage sum, 2 digits", func(t *testing.T) {
		input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
		var expected int = 357
		sum := GetSumMaxJoltage(input, 2)

		if sum != expected {
			t.Errorf("expected %d, but got %d", expected, sum)
		}
	},
	)
	t.Run("Should return maximum joltage sum, 12 digits", func(t *testing.T) {
		input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
		var expected int = 3121910778619
		sum := GetSumMaxJoltage(input, 12)

		if sum != expected {
			t.Errorf("expected %d, but got %d", expected, sum)
		}
	},
	)
}
