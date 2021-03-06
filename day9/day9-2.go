package main

// from https://www.jasoncoelho.com/2021/12/smoke-basin-advent-of-code-2021-day-9.html

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getCellVal(row int, col int, twoDArray [][]string) int {

	if row < 0 || col < 0 || row > len(twoDArray)-1 || col > len(twoDArray[row])-1 {
		return 10
	}

	var curCell, _ = strconv.Atoi(twoDArray[row][col])
	return curCell
}

func isLowPoint(row int, col int, twoDArray [][]string) bool {

	var curCell = getCellVal(row, col, twoDArray)

	// check neighbors to determine if this is a low point
	// only need to compare adjacent locations
	if curCell < getCellVal(row-1, col, twoDArray) &&
		curCell < getCellVal(row, col-1, twoDArray) && curCell < getCellVal(row, col+1, twoDArray) &&
		curCell < getCellVal(row+1, col, twoDArray) {
		return true
	}

	return false
}

func getBasinSize(row int, col int, twoDArray [][]string) int {

	var cellVal int = getCellVal(row, col, twoDArray)

	if cellVal >= 9 || cellVal == -1 {
		return 0
	}

	twoDArray[row][col] = strconv.Itoa(-1) // mark cell visited

	var northVal = getBasinSize(row-1, col, twoDArray)
	var southVal = getBasinSize(row+1, col, twoDArray)
	var eastVal = getBasinSize(row, col+1, twoDArray)
	var westVal = getBasinSize(row, col-1, twoDArray)

	return 1 + northVal + southVal + eastVal + westVal
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	var twoDArray [100][]string
	var lowPoints [500][3]int // oversized, i know
	var startRow = 0

	// get all the points into a 2D array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// could probably optimize and store int instead
		twoDArray[startRow] = strings.Split(scanner.Text(), "")
		startRow += 1
	}

	var lowPointCounter int = 0

	// first determine all the low points and the risk level sum
	for i := 0; i < 100; i++ {
		var row = twoDArray[i]
		if row != nil {

			for j := 0; j < len(row); j++ {

				if isLowPoint(i, j, twoDArray[:]) {
					var lowPoint = getCellVal(i, j, twoDArray[:])
					lowPoints[lowPointCounter] = [3]int{i, j, lowPoint} // store low points
					lowPointCounter += 1
				}
			}
		}
	}

	var basinSizes [300]int

	for i, lowPoint := range lowPoints {
		if lowPoint != [3]int{0, 0, 0} { // quit after we get "null" entries in this oversized array
			basinSizes[i] = getBasinSize(lowPoint[0], lowPoint[1], twoDArray[:])
		}
	}

	sort.Ints(basinSizes[:]) // cause we only need to top 3

	var solution2 int = basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]

	fmt.Println("Largest 3 basins multiplied ", solution2)
}
