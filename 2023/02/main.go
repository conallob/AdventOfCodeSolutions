package main

import (
        "bufio"
        "fmt"
	"regexp"
        "strings"
        "unicode"
)

var (
	MAX_CUBES = map[string]int{
		"blue": 14,
		"green": 13,
		"red": 12,
	} 
)

func GetGameNumber(line string) int {
	re := regexp.MustCompile(`Game (\d+):`)
	return int(re.FindString(line))
}

func IsGamePossible(line string) (int, bool) {
	gameNumber := GetGameNumber(line)
	 
	re := regexp.MustCompile(`\w(\d+)\w(red|blue|green)`)
	cubes := re.FindAllString(line) 
        if err != nil {
                log.Fatal(err)
        }
	for draw, colour := range cubes {
		if int(draw) > MAX_CUBES[colour] {
			return (gameNumber, false)
		}
	}
	return (gameNumber, true)
		
}

func main() {
        file, err := os.Open("exampletxt")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                fmt.Println(IsGamePossible(scanner.Text()))

        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        //fmt.Println("Final Total: ", totalCount)
}
