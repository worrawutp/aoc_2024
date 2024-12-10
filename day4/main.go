package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "slices"
	"strconv"
	"strings"
)

type Position struct {
    Row int
    Col int
}

var puzzleInput [][]string
var totalPositionFound [][][]int
var foundList []string
// var crossMasList []string
var crossMasCount int

func readPuzzle(filepath string) {
    file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for sc.Scan() {
        columns := strings.Split(sc.Text(), "")
        puzzleInput = append(puzzleInput, columns)
    }
}

func MaxCol() int {
    return len(puzzleInput[0])
}

func MaxRow() int {
    return len(puzzleInput)
}

func findXmasAt(direction string, p Position) {
    var r0, r1, r2, r3 int
    var c0, c1, c2, c3 int

    switch direction {
    case "R":
        if MaxCol() - p.Col >= 4 {
            r0 = p.Row; r1 = p.Row + 0; r2 = p.Row + 0; r3 = p.Row + 0
            c0 = p.Col; c1 = p.Col + 1; c2 = p.Col + 2; c3 = p.Col + 3
        } else {
            return
        }

    case "L":
        if p.Col >= 3 {
            r0 = p.Row; r1 = p.Row + 0; r2 = p.Row + 0; r3 = p.Row + 0
            c0 = p.Col; c1 = p.Col - 1; c2 = p.Col - 2; c3 = p.Col - 3
        } else {
            return
        }

    case "B":
        if MaxRow() - p.Row >=4 {
             r0 = p.Row; r1 = p.Row + 1; r2 = p.Row + 2; r3 = p.Row + 3
             c0 = p.Col; c1 = p.Col + 0; c2 = p.Col + 0; c3 = p.Col + 0
        } else {
            return
        }

    case "T":
        if p.Row >=3 {
            r0 = p.Row; r1 = p.Row - 1; r2 = p.Row - 2; r3 = p.Row - 3
            c0 = p.Col; c1 = p.Col + 0; c2 = p.Col + 0; c3 = p.Col + 0
        } else {
            return
        }

    case "BR":
        if (MaxRow() - p.Row >=4) && (MaxCol() - p.Col >= 4) {
            r0 = p.Row; r1 = p.Row + 1; r2 = p.Row + 2; r3 = p.Row + 3
            c0 = p.Col; c1 = p.Col + 1; c2 = p.Col + 2; c3 = p.Col + 3
        } else {
            return
        }

    case "BL":
        if (MaxRow() - p.Row >=4) && (p.Col >= 3) {
            r0 = p.Row; r1 = p.Row + 1; r2 = p.Row + 2; r3 = p.Row + 3
            c0 = p.Col; c1 = p.Col - 1; c2 = p.Col - 2; c3 = p.Col - 3
        } else {
            return
        }

    case "TR":
        if (p.Row >=3) && (MaxCol() - p.Col >= 4) {
            r0 = p.Row; r1 = p.Row - 1; r2 = p.Row - 2; r3 = p.Row - 3
            c0 = p.Col; c1 = p.Col + 1; c2 = p.Col + 2; c3 = p.Col + 3
        } else {
            return
        }

    case "TL":
        if (p.Row >=3) && (p.Col >= 3) {
            r0 = p.Row; r1 = p.Row - 1; r2 = p.Row - 2; r3 = p.Row - 3
            c0 = p.Col; c1 = p.Col - 1; c2 = p.Col - 2; c3 = p.Col - 3
        } else {
            return
        }
    }

    v := puzzleInput[r0][c0] + puzzleInput[r1][c1] + puzzleInput[r2][c2] + puzzleInput[r3][c3]
    if v == "XMAS" {
        matchPosition := [][]int{
            {r0, c0},
            {r1, c1},
            {r2, c2},
            {r3, c3},
        }
        // fmt.Printf("Bingo! %v %v\n", direction, matchPosition)
        updateFoundList(matchPosition)
    }
}

func matchCross(direction string, p Position) {
    switch direction {
    case "T":
        if puzzleInput[p.Row-1][p.Col-1] == "M" &&
            puzzleInput[p.Row-1][p.Col+1] == "M" &&
            puzzleInput[p.Row+1][p.Col-1] == "S" &&
            puzzleInput[p.Row+1][p.Col+1] == "S" {

            // fmt.Printf("match T | %v", p)
            crossMasCount += 1
            return
        }
    case "R":
        if puzzleInput[p.Row-1][p.Col-1] == "S" &&
            puzzleInput[p.Row-1][p.Col+1] == "M" &&
            puzzleInput[p.Row+1][p.Col-1] == "S" &&
            puzzleInput[p.Row+1][p.Col+1] == "M" {

            // fmt.Printf("match R | %v", p)
            crossMasCount += 1
            return
        }
     case "B":
         if puzzleInput[p.Row-1][p.Col-1] == "S" &&
             puzzleInput[p.Row-1][p.Col+1] == "S" &&
             puzzleInput[p.Row+1][p.Col-1] == "M" &&
             puzzleInput[p.Row+1][p.Col+1] == "M" {

             // fmt.Printf("match B | %v", p)
             crossMasCount += 1
             return
         }
    case "L":
        if puzzleInput[p.Row-1][p.Col-1] == "M" &&
            puzzleInput[p.Row-1][p.Col+1] == "S" &&
            puzzleInput[p.Row+1][p.Col-1] == "M" &&
            puzzleInput[p.Row+1][p.Col+1] == "S" {

            // fmt.Printf("match L | %v", p)
            crossMasCount += 1
            return
        }
    }
}

func findCrossMas(p Position, val string) {
    withinBorder := (p.Row >= 1 && p.Col >= 1 && p.Col < len(puzzleInput[0]) - 1 && p.Row < len(puzzleInput[0]) - 1)

    if val == "A" && withinBorder {
        matchCross("T", p)
        matchCross("R", p)
        matchCross("B", p)
        matchCross("L", p)
    }
}

func findXmas(p Position, val string) {
    // Only consider if it start with "X" or "S"
    if val == "X" || val == "S" {
        // fmt.Printf("%v | ", p)
        findXmasAt("R", p)
        findXmasAt("L", p)
        findXmasAt("B", p)
        findXmasAt("T", p)
        findXmasAt("BR", p)
        findXmasAt("BL", p)
        findXmasAt("TR", p)
        findXmasAt("TL", p)
    }
}

func positionStringReverse(position [][]int) string {
    result := ""
    for i := len(position) - 1; i>=0; i-- {
        result += strconv.Itoa(position[i][0]) + strconv.Itoa(position[i][1])
    } 
    return result
}

func positionString(position [][]int) string {
    result := ""
    for _, v := range position {
        result += strconv.Itoa(v[0]) + strconv.Itoa(v[1])
    } 
    return result
}

func updateFoundList(position [][]int) {
    if len(position) > 0 {
        foundList = append(foundList, positionString(position))
    }
}

func main() {
    totalPositionFound = [][][]int{}

    // readPuzzle("./input.txt") 
    // Example
    puzzleInput = [][]string{
        {"M","M","M","S","X","X","M","A","S","M"},
        {"M","S","A","M","X","M","S","M","S","A"},
        {"A","M","X","S","X","M","A","A","M","M"},
        {"M","S","A","M","A","S","M","S","M","X"},
        {"X","M","A","S","A","M","X","A","M","M"},
        {"X","X","A","M","M","X","X","A","M","A"},
        {"S","M","S","M","S","A","S","X","S","S"},
        {"S","A","X","A","M","A","S","A","A","A"},
        {"M","A","M","M","M","X","M","M","M","M"},
        {"M","X","M","X","A","X","M","A","S","X"},
    }

    for i, row := range puzzleInput {
        // fmt.Printf("\n\n Row #%v : %v\n", i, row)
        for j, col := range row {
            findXmas(Position{i, j}, col)
            findCrossMas(Position{i, j}, col)
        }
    }
    fmt.Println("------------ Result --------------")
    fmt.Printf("\n%v\n", len(foundList))
    fmt.Printf("\n%v\n", crossMasCount)
}
