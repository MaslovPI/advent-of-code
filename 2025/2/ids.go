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
	sum := GetSumInvalid(content)
	fmt.Print(sum)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetSumInvalid(ranges_str string) (sum int) {
	inputValues := strings.SplitSeq(ranges_str, ",")
	sum = 0

	for input := range inputValues {
		sum += getSumIds(GetAllInvalidInRange(input))
	}

	return
}

func getSumIds(invalidIds []int) (sum int) {
	sum = 0
	for _, id := range invalidIds {
		sum += id
	}
	return
}

func GetAllInvalidInRange(input string) []int {
	inputValues := strings.Split(strings.Trim(input, "\n"), "-")
	if len(inputValues) != 2 {
		panic("Invalid argument")
	}
	start, err := strconv.Atoi(inputValues[0])
	check(err)

	end, err := strconv.Atoi(inputValues[1])
	check(err)

	return getAllInvalidInRange(start, end)
}

func getAllInvalidInRange(start, end int) (result []int) {
	for i := start; i <= end; i++ {
		if IsInvalid(i) {
			result = append(result, i)
		}
	}
	return
}

func IsInvalid(id int) bool {
	idStr := strconv.Itoa(id)
	size := len(idStr)

	if size%2 != 0 {
		return false
	}

	middle := size / 2
	firstHalf := idStr[:middle]
	secondHalf := idStr[middle:]

	return firstHalf == secondHalf
}
