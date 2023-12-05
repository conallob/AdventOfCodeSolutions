package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
	"regexp"
	"strings"
	"strconv"
)

var (
	MAX_CUBES = map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	} 
)

func GetGameNumber(line string) string {
	re := regexp.MustCompile(`(\d+):`)
	return re.FindStringSubmatch(line)[1]
}

func IsGamePossible(line string) (string, bool) {
	gameNumber := GetGameNumber(line)
	 
	splitRegexp := regexp.MustCompile(`[\p{P}]`)
	cubes := splitRegexp.Split(line, -1)

	for _, draw := range cubes[1:] {
		drawResult := strings.Split(draw, " ")
		drawCount, err := strconv.Atoi(strings.Trim(drawResult[0], " "))
		if err != nil {
			log.Fatal(err)
		}
		if drawCount > MAX_CUBES[drawResult[1]] {
		  return gameNumber, false
		}
	}
	return gameNumber, true
		
}


func main() {
  file, err := os.Open("example.txt")
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
