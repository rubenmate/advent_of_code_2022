package main

import (
    "advent_of_code/day1"
	"os"
)

func main() {
    // day2.Run()
    if len(os.Args) == 1 {
        return
    }

    day := os.Args[1]

    switch day {
    case "1": 
        day1.Run()
    }
}
