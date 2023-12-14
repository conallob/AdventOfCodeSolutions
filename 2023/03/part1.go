package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func SearchAdjacent(x int, y int) bool {

	output := true

	for a, b := -1, -1 ; a < 2, b < 2 ; a++, b++ {
		if y - b >= 0 {
			 
		}	
	} 
}



func ScanForPartsNumbers(searchGrid *[][]rune, replace bool) rune {

	for y := 0 ; y < len(searchGrid) ; y++ {
		for x := 0 ; x < len(searchGrid[y]) ; x++ {
			if searchGrid[y][x] > 47 {
			}
		}
	}

}

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

	fmt.Println(ScanForPartsNumbers(schematic, true))
}
