package main

import (
	"adventOfCode/util"
	"errors"
	"fmt"
	"log"
	"os"
)

func part1(instructions string) int {
	floor := 0
	for _, char := range instructions {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	return floor
}

func part2(instructions string) (int, error) {
	floor := 0
	for instruction, char := range instructions {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor == -1 {
			return instruction + 1, nil
		}
	}
	return -1, errors.New("never reached basement")
}

func main() {
	file, err := os.Open("2015/day01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	stringSlice, err := util.ReadToString(file)

	if err != nil {
		log.Fatal(err)
	}

	resultPt1 := part1(stringSlice)
	resultPt2, err := part2(stringSlice)

	fmt.Println(resultPt1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resultPt2)
}
