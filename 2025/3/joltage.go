package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	content := string(dat)

	sum := GetSumMaxJoltage(content, 2)
	fmt.Printf("2 digits: %d\n", sum)

	sum = GetSumMaxJoltage(content, 12)
	fmt.Printf("12 digits: %d\n", sum)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetMaximumJoltage(input string, digits int) (result int) {
	length := len(input)
	result = 0
	id_from := 0
	for i := digits - 1; i >= 0; i-- {
		max, max_id := getMax(input, id_from, length-i)
		id_from = max_id + 1
		result += max * int(math.Pow(10.0, float64(i)))
	}
	return
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

func GetSumMaxJoltage(input string, digits int) (sum int) {
	banks := strings.Split(strings.TrimLeft(input, "\n"), "\n")
	sum = 0
	for _, bank := range banks {
		sum += GetMaximumJoltage(bank, digits)
	}
	return
}
