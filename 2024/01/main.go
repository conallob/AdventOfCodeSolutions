package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadandParseInput() (a []int, b []int) {
	// Open the file
	file, err := os.Open("2024/01/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		intA, _ := strconv.Atoi(line[0])
		intB, _ := strconv.Atoi(line[1])
		a = append(a, intA)
		b = append(b, intB)
	}

	return
}

func main() {
	listA, listB := ReadandParseInput()
	slices.Sort(listA)
	slices.Sort(listB)

	result := 0

	for i := range listA {
		result += abs(listA[i] - listB[i])
	}

	fmt.Println("Part 1: Sum of differences between the two lists is: ", result)

	similarityScore := 0
	for i := range listA {
		similarityScore += SearchSlice(listB, listA[i]) * listA[i]
	}

	fmt.Println("Part 2: Sum of similarity scores between the two lists is: ", similarityScore)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func SearchSlice(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}
