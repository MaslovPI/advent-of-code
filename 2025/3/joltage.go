package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	content := string(dat)
	sum := GetSumMaxJoltage(content)
	fmt.Printf("%d\n", sum)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetMaximumJoltage(input string) int {
	length := len(input)
	max_tens, max_tens_id := getMax(input, 0, length-1)
	max_singles, _ := getMax(input, max_tens_id+1, length)
	return max_tens*10 + max_singles
}

func getMax(input string, from, to int) (max int, id int) {
	max = 0
	id = -1

	for i := from; i < to; i++ {
		number := getNumber(input[i])

		if max < number {
			max = number
			id = i
		}

		if max == 9 {
			break
		}
	}
	return
}

var num_dict = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func getNumber(c byte) int {
	number := num_dict[c]
	return number
}

func GetSumMaxJoltage(input string) (sum int) {
	banks := strings.Split(strings.TrimLeft(input, "\n"), "\n")
	sum = 0
	for _, bank := range banks {
		sum += GetMaximumJoltage(bank)
	}
	return
}
