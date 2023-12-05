package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

// parse to find symbols (anything not a number or . )
type Point struct {
	X int
	Y int
}

func withinRange(a int, b int) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= 1
}

func main() {
	solve_p1()
	solve_p2()
}

func solve_p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := []Point{}
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, x := range line {
			if x != '.' && !unicode.IsDigit(x) {
				// fmt.Printf("x: %d, y: %d \n", i, lineNum)
				points = append(points, Point{X: i, Y: lineNum})
			}
		}
		lineNum += 1
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)
	lineNum = 0
	accum := 0

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		i := 0
		for i < len(runes) {
			if unicode.IsDigit(runes[i]) {
				// need to find a point in points within +/- 1 in both x and y
				found := false
				for _, point := range points {
					if withinRange(point.X, i) && withinRange(point.Y, lineNum) {
						found = true
						break
					}
				}
				if found {
					partNumber := ""
					y := 0
					for y := 1; i-y >= 0 && unicode.IsDigit(runes[i-y]); y++ {
						partNumber = string(runes[i-y]) + partNumber
					}

					partNumber += string(runes[i])

					for y = 1; i+y < len(runes) && unicode.IsDigit(runes[i+y]); y++ {
						partNumber += string(runes[i+y])
					}
					i += y

					num, err := strconv.Atoi(partNumber)
					if err != nil {
						log.Fatal("Error during conversion:", err)
					}
					// fmt.Println(num)
					accum += num
				} else {
					i++
				}
			} else {
				i++
			}
		}
		lineNum++
	}
	fmt.Println("Part one: ", accum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solve_p2() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	dx := []int{-1, 0, 1, 0, -1, -1, 1, 1} // 8 directions: up, right, down, left and diagonals
	dy := []int{0, 1, 0, -1, -1, 1, 1, -1}
	n, m := len(grid), len(grid[0])
	totalSum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' {
				gears := []int{}

				// Check at most two neighbouring gears.
				for k := 0; k < 8 && len(gears) < 2; k++ {
					ni, nj := i+dx[k], j+dy[k]
					if ni >= 0 && ni < n && nj >= 0 && nj < m && unicode.IsDigit(grid[ni][nj]) {
						partNumber := ""
						// Extend to the left and right for multi-digit numbers
						l, r := nj, nj
						for ; l-1 >= 0 && unicode.IsDigit(grid[ni][l-1]); l-- {
						}
						for ; r+1 < m && unicode.IsDigit(grid[ni][r+1]); r++ {
						}
						partNumber = string(grid[ni][l : r+1])
						gear, _ := strconv.Atoi(partNumber)
						gears = append(gears, gear)
					}
				}

				if len(gears) == 2 {
					totalSum += gears[0] * gears[1]
				}
			}
		}
	}

	fmt.Println("Total sum:", totalSum)
}

//func solve_p2() {
//file, err := os.Open("test.txt")
//if err != nil {
//log.Fatal(err)
//}
//defer file.Close()

//points := []Point{}
//scanner := bufio.NewScanner(file)
//lineNum := 0
//for scanner.Scan() {
////line := scanner.Text()
//for i, x := range line {
//if x == '*' {
////fmt.Printf("x: %d, y: %d \n", i, lineNum)
//points = append(points, Point{X: i, Y: lineNum})
//}
//}
//lineNum += 1
//}
//_, err = file.Seek(0, io.SeekStart)
//if err != nil {
//log.Fatal(err)
//}
//scanner = bufio.NewScanner(file)
//lineNum = 0
//accum := 0
//firstGear := 0
//firstGearAt := Point{X: -1, Y: -1}
//for scanner.Scan() {
////line := scanner.Text()
//runes := []rune(line)
//i := 0
//for i < len(runes) {
//if unicode.IsDigit(runes[i]) {
////need to find a point in points within +/- 1 in both x and y
//found := false
//workingPoint := Point{X: -1, Y: -1}
//for _, point := range points {
//if withinRange(point.X, i) && withinRange(point.Y, lineNum) {
//found = true
//workingPoint = point
//break
//}
//}
//if found {
//partNumber := ""
//y := 0
//for y := 1; i-y >= 0 && unicode.IsDigit(runes[i-y]); y++ {
//partNumber = string(runes[i-y]) + partNumber
//}

//partNumber += string(runes[i])

//for y = 1; i+y < len(runes) && unicode.IsDigit(runes[i+y]); y++ {
//partNumber += string(runes[i+y])
//}

//num, err := strconv.Atoi(partNumber)
//if err != nil {
//log.Fatal("Error during conversion:", err)
//}

//if firstGear != 0 {
//if withinRange(workingPoint.X, firstGearAt.X) && withinRange(workingPoint.Y, firstGearAt.Y) {
//accum += (firstGear * num)
//firstGear = 0
//firstGearAt = Point{X: -1, Y: -1}
//} else {
//firstGearAt = Point{X: i, Y: lineNum}
//firstGear = num
//}
//} else {
//firstGearAt = Point{X: i, Y: lineNum}
//firstGear = num
//}

//i += y
//} else {
//i++
//}
//} else {
//i++
//}
//}
//lineNum++
//}

//fmt.Println("Part two: ", accum)

//if err := scanner.Err(); err != nil {
//log.Fatal(err)
//}
//}
