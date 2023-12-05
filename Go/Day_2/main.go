package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	accum := 0
	//red := 12
	//green := 13
	//blue := 14

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split_one := strings.Split(line, "Game ")
		split_two := strings.Split(split_one[1], ":")
		//game := split_two[0]
		//num, err := strconv.Atoi(game)
		//if err != nil {
		//fmt.Println(err)
		//}
		//fmt.Println(num)
		inst_split := strings.Split(split_two[1], ";")
		//good := true
		lowest_red := 0
		lowest_green := 0
		lowest_blue := 0
		for _, game := range inst_split {
			//fmt.Println(i)
			for _, pull := range strings.Split(game, ",") {
				if strings.Contains(pull, "red") {
					split_for_num := strings.Split(pull, " ")
					count, err := strconv.Atoi(split_for_num[1])
					if err != nil {
						fmt.Println(err)
					}
					if count > lowest_red {
						lowest_red = count
					}
				}
				if strings.Contains(pull, "green") {
					split_for_num := strings.Split(pull, " ")
					count, err := strconv.Atoi(split_for_num[1])
					if err != nil {
						fmt.Println(err)
					}
					if count > lowest_green {
						lowest_green = count
					}

				}
				if strings.Contains(pull, "blue") {
					split_for_num := strings.Split(pull, " ")
					count, err := strconv.Atoi(split_for_num[1])
					if err != nil {
						fmt.Println(err)
					}
					if count > lowest_blue {
						lowest_blue = count
					}
				}
			}
		}
		accum += (lowest_blue * lowest_green * lowest_red)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(accum)
}
