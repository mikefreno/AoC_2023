package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringToIntList(str string) []int {
	cleanedPrefix := strings.TrimPrefix(str, " ")
	cleanedSuffix := strings.TrimSuffix(cleanedPrefix, " ")
	numsString := strings.Split(cleanedSuffix, " ")
	var numList []int
	for _, x := range numsString {
		if x != "" {
			num, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			numList = append(numList, num)
		}
	}
	return numList
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cardCounts := make([]int, 1000)
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	problemOne := 0
	problemTwo := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstSplit := strings.Split(line, ": ")
		numSets := strings.Split(firstSplit[1], "|")
		targets := stringToIntList(numSets[0])
		pulls := stringToIntList(numSets[1])
		matches := 0
		for _, n := range pulls {
			for _, x := range targets {
				if n == x {
					matches++
				}
			}
		}
		n := float64(matches)
		problemOne += int(1 * math.Pow(2, n-1))
		numberOfCurrentCard := cardCounts[0]
		problemTwo += numberOfCurrentCard
		for i := 0; i < matches; i++ {
			cardCounts[i+1] += numberOfCurrentCard
		}
		cardCounts = cardCounts[1:]

	}
	fmt.Println("problem 1: ", problemOne)
	fmt.Println("problem 2: ", problemTwo)
}
