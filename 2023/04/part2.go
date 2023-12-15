package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"

	"github.com/juliangruber/go-intersect"
)


func FindAllWinningCards(start int) 

func ParseCardNum(card string) int {
	out := strings.TrimSuffix(strings.TrimPrefix(card, "Card "), ":")
	outAsInt, _ := strconv.Atoi(strings.TrimSpace(out))
	return outAsInt
}

func SanitizeWinningScratchcard(nums []interface{}) ([]int, error) {
	out := []int{}

	// Filter out nil values
	for _, a := range nums {
		if ((a == "") || (a == nil) || (a == " ")) {
			continue
		} else {
			AAsInt, _ := strconv.Atoi(fmt.Sprintf("%v", a))
			out = append(out, AAsInt)
		}
	}

	if len(out) <= 0 {
		return nil, fmt.Errorf("Not a winner")
	} else {
		return out, nil
	}
}


func main() {
	file, err := os.Open("example-part1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
 	re := regexp.MustCompile(`[[:punct:]]`)
	ScrtchCardSet := map[int][]int{}

	for scanner.Scan() {
  	parsed := re.Split(scanner.Text(), -1)

		card := ParseCardNum(parsed[0])
		winners := strings.Split(parsed[1], " ")
		scratchcard := strings.Split(parsed[2], " ")
		winningNums := intersect.Simple(winners, scratchcard)

		ScrtchCardSet[card], _ = SanitizeWinningScratchcard(winningNums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalCards := len(ScrtchCardSet)

	for _, game := range ScratchCardSet ; game != [] {
		totalCards += len(
		
	}
	

}
