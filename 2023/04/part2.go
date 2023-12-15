package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/juliangruber/go-intersect"
)


func ParseCardNum(card string) string {
	out := strings.TrimSuffix(strings.TrimPrefix(card, "Card "), ":")
	return strings.TrimSpace(out)
}

func SanitizeWinningScratchcard(nums []interface{}) ([]string, error) {

	out := []string{}

	// Filter out nil values
	for _, a := range nums {
		if ((a == "") || (a == nil) || (a == " ")) {
			continue
		} else {
			// TODO(conallob): Convert to int, allowing to walk a winning chain
			out = append(out, fmt.Sprintf("%v", a))
		}
	}

	if len(out) <= 0 {
		return nil, fmt.Errorf("Not a winner")
	} else {
		return out, nil
	}

}


func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
 	re := regexp.MustCompile(`[[:punct:]]`)
	ScrtchCardSet := map[string][]string{}

	for scanner.Scan() {
  	parsed := re.Split(scanner.Text(), -1)

		card := ParseCardNum(parsed[0])
		winners := strings.Split(parsed[1], " ")
		scratchcard := strings.Split(parsed[2], " ")
		winningNums := intersect.Simple(winners, scratchcard)

		ScrtchCardSet[card], _ = SanitizeWinningScratchcard(winningNums)


	}

	fmt.Println(ScrtchCardSet)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
