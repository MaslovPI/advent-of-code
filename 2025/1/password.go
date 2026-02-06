package main

import (
	"fmt"
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
	fmt.Print(password)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CalculatePassword(moves []string) int {
	currentPosition := 50
	password := 0
	for _, move := range moves {
		if move == "" {
			continue
		}

		currentPosition = compensate(adjust(currentPosition, move))
		if currentPosition == 0 {
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

func compensate(position int) int {
	for position < 0 {
		position += 100
	}

	for position > 99 {
		position -= 100
	}

	return position
}
