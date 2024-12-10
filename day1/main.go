package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func smallest(s []int) int {
    return slices.Min(s)
}

func Abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func removeWithMap(slice []int, value int) []int {
    result := []int{}
    for _, v := range slice {
        if v != value {
            result = append(result, v)
        }
    }
    return result
}

func RemoveOnce(slice []int, value int) []int {
    result := slice[:0]
    picked := false
    for _, v := range slice {
        if v == value && picked == false {
            picked = true
        } else {
            result = append(result, v)
        }
    }
    return result
}

func SumValue(items []int) int {
    sum := 0
    for _, d := range items {
        sum += d
    }
    return sum
}

func SimilarityCount(v int, items []int) int {
    result := 0
    for _, item := range items {
        if v == item {
            result += 1
        }
    }
    return result
}

func main() {
    f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatalf("open file error: %v", err)
        return
    }
    defer f.Close() 

    left := []int{}
    right := []int{}

    sc := bufio.NewScanner(f)
    for sc.Scan() {
        line := strings.Split(sc.Text(), "   ")
        leftValue, _ := strconv.Atoi(line[0])
        rightValue, _ := strconv.Atoi(line[1])
        left = append(left, leftValue)
        right = append(right, rightValue)
    }

    // find distances
    distances := []int{}
    dLeft := make([]int, len(left))
    dRight := make([]int, len(right))
    copy(dLeft, left)
    copy(dRight, right)
    for len(dLeft) > 0 && len(dRight) > 0 {
        smallestLeft := smallest(dLeft)
        smallestRight := smallest(dRight)
        distance := Abs(smallestLeft - smallestRight)
        distances = append(distances, distance)

        // Delete the smallestLeft and smallestRight from the slices
        dLeft = RemoveOnce(dLeft, smallestLeft)
        dRight = RemoveOnce(dRight, smallestRight)
    }
    fmt.Println("------- Total Distances ---------")
    fmt.Printf("%v\n", SumValue(distances))

    // find similarities
    similarities := []int{}
    sLeft := make([]int, len(left))
    sRight := make([]int, len(right))
    copy(sLeft, left)
    copy(sRight, right)
    for _, value := range sLeft {
        count := SimilarityCount(value, sRight)
        score := value * count
        similarities = append(similarities, score)
    }
    fmt.Println("------- Total Similarity Score ---------")
    fmt.Printf("%v\n", SumValue(similarities))
}
