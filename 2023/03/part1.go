package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example-part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	schematic := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schematic = append(schematic, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ScanForPartsNumbers(schematic))
}
