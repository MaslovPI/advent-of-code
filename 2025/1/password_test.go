package main

import "testing"

func TestCalculatePassword(t *testing.T) {
	t.Run("Should calculate correct password", func(t *testing.T) {
		moves := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		expected := 3
		password := CalculatePassword(moves)

		if password != expected {
			t.Errorf("expected %d, but got %d", expected, password)
		}
	})
	t.Run("Should correctly overcompensate", func(t *testing.T) {
		moves := []string{
			"L350",
			"R900",
		}
		expected := 2
		password := CalculatePassword(moves)

		if password != expected {
			t.Errorf("expected %d, but got %d", expected, password)
		}
	})
}

func TestCalculateNewPassword(t *testing.T) {
	t.Run("Should calculate correct password", func(t *testing.T) {
		moves := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		expected := 6
		password := CalculateNewPassword(moves)

		if password != expected {
			t.Errorf("expected %d, but got %d", expected, password)
		}
	})
	t.Run("Should correctly overcompensate", func(t *testing.T) {
		moves := []string{
			"L350",
			"R900",
		}
		expected := 13
		password := CalculateNewPassword(moves)

		if password != expected {
			t.Errorf("expected %d, but got %d", expected, password)
		}
	})
}
