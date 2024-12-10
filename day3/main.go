package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
    "strconv"
)

func readData(filepath string) ([]string, error) {
    result := []string{}
    file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm) 
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for sc.Scan() {
        result = append(result, sc.Text())
    }
    return result, nil
}

func buildMuliplyList(data string) []int {
    r, _ := regexp.Compile(`mul\((\d{1,3}\,\d{1,3})\)`)
    matchData := r.FindAllStringSubmatch(data, -1)

    result := []int{}
    for _, v := range matchData {
        params := strings.Split(v[1], ",")
        d1, _ := strconv.Atoi(params[0])
        d2, _ := strconv.Atoi(params[1])
        result = append(result, d1 * d2)
    }
    return result
}

func main() {
    allInput, _ := readData("./input.txt") 
    var sum int64
    for _, input := range allInput {
        listNum := []int{}
        listNum = buildMuliplyList(input)
        for _, v := range listNum {
            sum += int64(v)
        }
    } 
    fmt.Println("----------- Result ------------")
    fmt.Printf("%v\n", sum)
}
