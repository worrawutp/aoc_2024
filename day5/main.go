package main

import "fmt"

func loadPuzzle(filepath string) (rules [][]int, updateList []int) {


    return
}

func main() {
    rules := [][]int{}
    updateList := []int{}

    rules, updateList = loadPuzzle("./input.txt")

    fmt.Printf("%v\n", rules)
    fmt.Printf("%v\n", updateList)
}
