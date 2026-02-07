package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	content := string(dat)
	moves := strings.Split(content, "\n")
	password := CalculatePassword(moves)
	fmt.Printf("%d\n", password)
	newPassword := CalculateNewPassword(moves)
	fmt.Printf("%d\n", newPassword)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CalculatePassword(moves []string) int {
	currentPosition := 50
	password := 0
	for _, move := range moves {
		if move == "" {
			continue
		}

		currentPosition, _ = compensate(adjust(currentPosition, move), false)
		if currentPosition == 0 {
			password++
		}
	}
	return password
}

func CalculateNewPassword(moves []string) int {
	currentPosition := 50
	password := 0
	for _, move := range moves {
		if move == "" {
			continue
		}

		startedAtZero := currentPosition == 0
		pos, compensations := compensate(adjust(currentPosition, move), startedAtZero)
		currentPosition = pos

		if compensations > 0 {
			password += compensations
		} else if currentPosition == 0 {
			password++
		}
	}
	return password
}

func adjust(position int, move string) int {
	letter := move[0]
	num_str := move[1:]

	num, err := strconv.Atoi(num_str)
	check(err)

	if letter == 'R' {
		return position + num
	}
	return position - num
}

func compensate(position int, startedAtZero bool) (int, int) {
	compensations := 0
	isBelowZero := position < 0
	for position < 0 {
		position += 100
		compensations++
	}

	for position > 99 {
		position -= 100
		compensations++
	}

	if isBelowZero && compensations > 0 {
		if position == 0 {
			compensations++
		}
		if startedAtZero {
			compensations--
		}
	}

	return position, compensations
}
