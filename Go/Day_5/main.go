package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type coverMap struct {
	DestinationStart int
	SourceStart      int
	RangeLength      int
}
type seedRange struct {
	Start int
	Range int
}

func withinRange(sourceStart int, rangeLength int, num int) bool {
	if num < sourceStart {
		return false
	}
	diff := num - sourceStart
	if diff < rangeLength {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seeds []seedRange
	var seedToSoil []coverMap
	var soilToFert []coverMap
	var fertToWater []coverMap
	var waterToLight []coverMap
	var lightToTemp []coverMap
	var tempToHumidity []coverMap
	var humidityToLoc []coverMap

	readingStep := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingStep++
			continue
		}
		switch readingStep {
		case 0:
			//reading seeds
			numsString := strings.Split(line, "seeds: ")[1]
			nums := strings.Split(numsString, " ")

			start := -1
			for _, n := range nums {
				num, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal("Error during conversion:", err)
				}
				if start == -1 {
					start = num
				} else {
					seeds = append(seeds, seedRange{Start: start, Range: num})
					start = -1
				}

			}
		case 1:
			//reading seed to soil
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			seedToSoil = append(seedToSoil, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		case 2:
			//reading soil to fert
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			soilToFert = append(soilToFert, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		case 3:
			//reading fert to water
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			fertToWater = append(fertToWater, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		case 4:
			//reading water to light
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			waterToLight = append(waterToLight, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		case 5:
			//reading light to temp
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			lightToTemp = append(lightToTemp, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		case 6:
			//reading temp to hum
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			tempToHumidity = append(tempToHumidity, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		case 7:
			//reading hum to loc
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			split := strings.Split(line, " ")
			dest := split[0]
			destNum, err := strconv.Atoi(dest)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			source := split[1]
			sourceNum, err := strconv.Atoi(source)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			rangeStart := split[2]
			rangeNum, err := strconv.Atoi(rangeStart)
			if err != nil {
				log.Fatal("Error during conversion:", err)
			}
			humidityToLoc = append(humidityToLoc, coverMap{DestinationStart: destNum, SourceStart: sourceNum, RangeLength: rangeNum})
		}
	}

	min := 10000000000
	for _, x := range seeds {
		for i := x.Start; i < x.Start+x.Range; i++ {
			seed := i
			for _, seedSoilLine := range seedToSoil {
				if withinRange(seedSoilLine.SourceStart, seedSoilLine.RangeLength, seed) {
					distance := seed - seedSoilLine.SourceStart
					seed = seedSoilLine.DestinationStart + distance
					break
				}
			}

			for _, soilToFertLine := range soilToFert {
				if withinRange(soilToFertLine.SourceStart, soilToFertLine.RangeLength, seed) {
					distance := seed - soilToFertLine.SourceStart
					seed = soilToFertLine.DestinationStart + distance
					break
				}
			}

			for _, fertToWaterLine := range fertToWater {
				if withinRange(fertToWaterLine.SourceStart, fertToWaterLine.RangeLength, seed) {
					distance := seed - fertToWaterLine.SourceStart
					seed = fertToWaterLine.DestinationStart + distance
					break
				}
			}

			for _, waterToLightLine := range waterToLight {
				if withinRange(waterToLightLine.SourceStart, waterToLightLine.RangeLength, seed) {
					distance := seed - waterToLightLine.SourceStart
					seed = waterToLightLine.DestinationStart + distance
					break
				}
			}

			for _, lightToTempLine := range lightToTemp {
				if withinRange(lightToTempLine.SourceStart, lightToTempLine.RangeLength, seed) {
					distance := seed - lightToTempLine.SourceStart
					seed = lightToTempLine.DestinationStart + distance
					break
				}
			}

			for _, tempToHumidity := range tempToHumidity {
				if withinRange(tempToHumidity.SourceStart, tempToHumidity.RangeLength, seed) {
					distance := seed - tempToHumidity.SourceStart
					seed = tempToHumidity.DestinationStart + distance
					break
				}
			}

			for _, humidityToLoc := range humidityToLoc {
				if withinRange(humidityToLoc.SourceStart, humidityToLoc.RangeLength, seed) {
					distance := seed - humidityToLoc.SourceStart
					seed = humidityToLoc.DestinationStart + distance
					break
				}
			}

			if seed < min {
				min = seed
			}
		}
	}
	fmt.Println(min)
	elapsed := time.Since(start)
	fmt.Printf("The code took %s to execute.\n", elapsed)

}
