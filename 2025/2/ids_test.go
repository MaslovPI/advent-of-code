package main

import (
	"reflect"
	"testing"
)

func TestGetSumInvalid(t *testing.T) {
	t.Run("Should sum all invalid ids in ranges", func(t *testing.T) {
		ranges_str := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124\n"
		expected := 1227775554
		sum := GetSumInvalid(ranges_str)

		if sum != expected {
			t.Errorf("expected %d, but got %d", expected, sum)
		}
	},
	)
}

func TestGetAllInvalidInRange(t *testing.T) {
	t.Run("Should find all invalid ids in given range", func(t *testing.T) {
		test_cases := map[string][]int{
			"11-22":                 {11, 22},
			"95-115":                {99},
			"998-1012":              {1010},
			"1188511880-1188511890": {1188511885},
			"222220-222224":         {222222},
			"1698522-1698528":       nil,
			"446443-446449":         {446446},
			"38593856-38593862":     {38593859},
			"565653-565659":         nil,
			"824824821-824824827":   nil,
			"2121212118-2121212124": nil,
		}

		for input, expected := range test_cases {
			result := GetAllInvalidInRange(input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("expected %#v, but got %#v", expected, result)
			}
		}
	})
}

func TestIsInvalid(t *testing.T) {
	t.Run("Should detect invalid ids", func(t *testing.T) {
		test_cases := []int{
			11,
			22,
			99,
			1010,
			1188511885,
			222222,
			446446,
			38593859,
		}

		for _, id := range test_cases {
			result := IsInvalid(id)
			if !result {
				t.Errorf("expected %d to be invalid", id)
			}
		}
	})
	t.Run("Should detect valid ids", func(t *testing.T) {
		test_cases := []int{
			12,
			25,
			999,
			1011,
			121212,
			81,
			0,
		}

		for _, id := range test_cases {
			result := IsInvalid(id)
			if result {
				t.Errorf("expected %d to be valid", id)
			}
		}
	})
}
