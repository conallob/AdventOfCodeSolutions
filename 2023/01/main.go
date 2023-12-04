package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func PrepareNumbers(input string) string {
	numericWords := map[int]string{
	 	9:	"nine",
	 	8:	"eight",
	 	7:	"seven",
	 	6:	"six",
	 	5:	"five",
	 	4:	"four",
	 	3:	"three",
	 	2:	"two",
	 	1:	"one",
	}
	out := input

	ordering := []int{2, 8, 4, 1, 3, 9, 5, 6, 7}
	
	for _, i := range ordering {
		out = strings.Replace(out, numericWords[i], string(rune(i) + 48), 2)
	}
	return out
}

func FindAllNumbers(input string) []int {
	var output []int

	preparedInput := PrepareNumbers(input)
	runeInput := []rune(preparedInput)

	if strings.ContainsAny(preparedInput, "1234567890") {
		for _, x := range runeInput {
			if unicode.IsDigit(x) {
				// Convert from ASCII table entry into actual digit
				output = append(output, (int(x) - 48))
			}
		}
	} else {
		output = append(output, 0)
	}
	return output
}

func main() {
	file, err := os.Open("example-part2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := FindAllNumbers(scanner.Text())
		fmt.Println(scanner.Text(), nums)
		switch {
		case len(nums) == 0:
			continue
		case len(nums) == 1:
			totalCount = totalCount + (nums[0] * 10) + (nums[0])
		case len(nums) > 1:
			totalCount = totalCount + (nums[0] * 10) + (nums[len(nums)-1])
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final Total: ", totalCount)
}
