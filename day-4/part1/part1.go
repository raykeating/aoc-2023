package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

	var sum int = 0
	var scratchCards []map[string][]int
	
	var i = 0
    for scanner.Scan() {
        var line string = scanner.Text()
        fmt.Println(line)

		line = strings.Split(line, ": ")[1]
		var lhs string = strings.Split(line, " | ")[0]
		var rhs string = strings.Split(line, " | ")[1]
		var winningNumsStrings []string = strings.Split(lhs, " ")
		var ticketNumsStrings []string = strings.Split(rhs, " ")

		var cards map[string][]int
		cards = make(map[string][]int)

		for _, v := range winningNumsStrings {
			var num, _ = strconv.Atoi(v)
			if (num != 0) {
				cards["winning"] = append(cards["winning"], num)
			}
		}

		for _, v := range ticketNumsStrings {
			var num, _ = strconv.Atoi(v)
			if (num != 0) {
				cards["ticket"] = append(cards["ticket"], num)
			}
		}

		scratchCards = append(scratchCards, cards)
		i++
    }



	for _, cards := range scratchCards {

		var winningNumCount = 0

		for _, winningNum := range cards["winning"] {
			if (slices.Contains(cards["ticket"], winningNum)) {
				winningNumCount += 1
			}
		}

		sum += getPoints(winningNumCount)

	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
}

func getPoints(x int) int {
	if x == 1 {
		return x
	} else {
		return int(math.Pow(float64(2), float64(x-1)))
	}
}