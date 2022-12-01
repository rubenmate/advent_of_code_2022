package day1

import (
	// "advent_of_code/helpers"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const path = "day1/input.txt"

func Run() {
    input,_ := os.Open(path)
	defer input.Close()
    scanner := bufio.NewScanner(input)

    currentCalories := 0
    var top3Elfs = []int{-3, -2, -1} // init
    for scanner.Scan() {
        food, err := strconv.Atoi(scanner.Text())
        if err != nil {
            if currentCalories > top3Elfs[0]{
                top3Elfs[0] = currentCalories
                elfsSlice := top3Elfs[:]
                sort.Ints(elfsSlice)
                top3Elfs = elfsSlice
            }
            currentCalories = 0
        } else {
            currentCalories += food
        }
    }

    fmt.Println("-Part 1: Max calories elf has", getMax(top3Elfs))
    fmt.Println("-Part 2: Top 3 elfs are", top3Elfs, "and the sum is", sum(top3Elfs))
}

func getMax(array []int) (max int) {
    max = 0
    for _, v := range array {
        if v > max {
            max = v
        }
    }
    return max
}

func sum(array []int) (sum int) {
    sum = 0
    for _, elf := range array{
        sum += elf
    }
    return sum
}


