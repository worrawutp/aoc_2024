package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func ReadReport(filepath string) ([][]int, error) {
    result := [][]int{}
    file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm) 
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for sc.Scan() {
        line := strings.Split(sc.Text(), " ")  
        levels := []int{}
        for _, v := range line {
            lv, _ := strconv.Atoi(v)
            levels = append(levels, lv)
        }
        result = append(result, levels)
    }

    return result, nil
}

func setDirection(increasing bool) (direction string)  {
    if increasing {
        direction = "increasing"
    } else {
        direction = "decreasing"
    }
    return
}

func validateLevels(levels []int) bool {
    unsafe := false
    direction := ""
    previousDirection := ""

    for j, _ := range levels {
        // Skip the last level
        if j == (len(levels) - 1) {
           continue 
        }
        nextIndex := j + 1

        increasing := levels[j] < levels[nextIndex]

        // 1st iteration only
        if previousDirection == "" && direction == "" {
            direction = setDirection(increasing)
        }

        previousDirection = direction
        direction = setDirection(increasing)

        levelDiff := Abs(levels[j] - levels[nextIndex])
        isNotCorrectLevelDiff := levelDiff < 1 || levelDiff > 3
        isNotSameDirection := previousDirection != direction 
        if isNotSameDirection || isNotCorrectLevelDiff {
            unsafe = true
            break
        }
    } // levels

    return !unsafe
}

func removeSliceByIndex(slice []int, i int) []int {
    return append(slice[:i], slice[i+1:]...)
}

func tolerateSingleBadLevel(levels []int) bool {
    // fmt.Println("Check tolerateSingleBadLevel...")
    // fmt.Println(levels)
    // Try remove each item and then check if levels still good
    for i := 0; i < len(levels); i++ {
        newSlice := make([]int, 0, len(levels)-1)
        newSlice = append(newSlice, levels[:i]...)
        newSlice = append(newSlice, levels[i+1:]...)
        valid := validateLevels(newSlice)

        // fmt.Printf("removing %v | %v | %v\n", i, newSlice ,valid)
        if valid {
            return true
        }
    }
    return false
}

func main() {
    result := map[string]int{
        "safe": 0,
        "unsafe": 0,
    }

    reports, _ := ReadReport("./input.txt")

    for _, levels := range reports {
        valid := validateLevels(levels)
        if valid {
            result["safe"] += 1
        } else {
            // check if it can "tolerate a single bad level"
            if tolerateSingleBadLevel(levels) {
                result["safe"] += 1
            } else {
                result["unsafe"] += 1
            }
        }
    }

    fmt.Println("---------- Result ------------")
    fmt.Printf("%v\n", result)
}
