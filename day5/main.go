package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func loadPuzzle(filepath string) (rules [][]string, updateList [][]string) {
    file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
    if err != nil {
        panic("Cannot open puzzle file!")
    }
    
    sc := bufio.NewScanner(file)
    readRule := true
    for sc.Scan() {
        line := sc.Text()
        if line == "" {
            readRule = false
        }

        if readRule {
            rules = append(rules, strings.Split(line, "|"))
        } else {
            if(len(line) > 0) {
                updateList = append(updateList, strings.Split(line, ","))
            }
        }
    }
    return
}

func loadExample() (rules [][]string, updateList [][]string){
    rulesStr := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`
    lines := strings.Split(rulesStr, "\n")
    for _, line := range lines {
        rules = append(rules, strings.Split(line, "|"))
    }

    aaa := `75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
    lines = strings.Split(aaa, "\n")
    for _, line := range lines {
        updateList = append(updateList, strings.Split(line, ","))
    }

    return
}

func ruleSetNotInTheList(list []string, rule []string) bool {
    for _, r := range rule {
        if !slices.Contains(list, r) {
            return true 
        }
    } 
    return false
}

func Qualified(list []string, rule []string) bool {
    iX := 0
    iY := 0
    for i, v := range list {
        if v == rule[0] {
            iX = i
        }
        if v == rule[1] {
            iY = i
        }
    }

    if iX < iY {
        // X|Y, X come before Y, thus the Right order
        return true
    } else {
        return false
    }
}

func SumMiddle(allList [][]string) int {
    // fmt.Println("Calculate SumMiddle...")
    // fmt.Println(qualifiedList)
    result := 0

    for _, list := range allList {
        middleIndex := (len(list) - 1) / 2
        middleValue, err := strconv.Atoi(list[middleIndex])
        if err != nil {
            panic("Something wrong during calculate SumMiddle!")
        }
        result += middleValue
    }
    return result
}

func sortOutOrderInList(list []string, rules [][]string) []string {
    iX := 0
    iY := 0
    for _, rule := range rules {
        if ruleSetNotInTheList(list, rule) { continue }
        if Qualified(list, rule) { continue }

        for i, v := range list {
            if v == rule[0] {
                iX = i
            }
            if v == rule[1] {
                iY = i
            }

            if iX > iY {
                // swap
                list[iX] = rule[1]
                list[iY] = rule[0]
                // fmt.Printf("--> %v |  %v\n", rule, list)
                return sortOutOrderInList(list, rules) 
            }
        }
    }
    return list
}

func reBuildBadList(badLists [][]string, rules [][]string) [][]string {
    result := [][]string{}
    aaa := []string{}

    for _, list := range badLists {
        fmt.Println(list)
        aaa = sortOutOrderInList(list, rules)
        result = append(result, aaa)
    }

    return result
}

func main() {
    rules := [][]string{}
    inputList := [][]string{}
    qualifiedList := [][]string{}
    badList := [][]string{}
    fixedList := [][]string{}

    rules, inputList = loadPuzzle("./input.txt")
    // rules, inputList = loadExample()

    for _, list := range inputList {
        // fmt.Println("Next list > %v", list)
        goodList := true

        for _, rule := range rules {
            if ruleSetNotInTheList(list, rule) {
                // fmt.Printf("Rule not in the list : %v\n", rule)
                continue
            }
            if !Qualified(list, rule) {
                badList = append(badList, list)
                goodList = false
                break
            }
        }

        // all rules are verified
        // add this list into the qualified list
        if goodList && len(list) > 0 {
            // fmt.Printf("%v is qualified\n", list)
            qualifiedList = append(qualifiedList, list)
        }
    }
    result := SumMiddle(qualifiedList)

    fixedList = reBuildBadList(badList, rules)
    result2 := SumMiddle(fixedList)

    // fmt.Printf("%v\n", rules)
    // fmt.Printf("%v\n", inputList)
    fmt.Printf("\n--------------\n")
    // fmt.Printf("%v\n", qualifiedList)
    fmt.Println(result)

    // fmt.Printf("%v\n", fixedList)
    fmt.Println(result2)
}
