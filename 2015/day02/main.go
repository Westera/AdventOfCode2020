package main

import (
	"adventOfCode/util"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(dimensions []string) int {
	sum := 0
	for _, dimension := range dimensions {
		xyz := strings.Split(dimension, "x")
		var intXYZ = []int{}
		for _, i := range xyz {
			j, _ := strconv.Atoi(i)
			intXYZ = append(intXYZ, j)
		}
		sort.Ints(intXYZ)
		value := 3*intXYZ[0]*intXYZ[1] + 2*intXYZ[1]*intXYZ[2] + 2*intXYZ[0]*intXYZ[2]
		sum += value
	}

	return sum
}

func part2(dimensions []string) int {
	sum := 0
	for _, dimension := range dimensions {
		xyz := strings.Split(dimension, "x")
		var intXYZ = []int{}
		for _, i := range xyz {
			j, _ := strconv.Atoi(i)
			intXYZ = append(intXYZ, j)
		}
		sort.Ints(intXYZ)
		value := 2*intXYZ[0] + 2*intXYZ[1] + intXYZ[0]*intXYZ[1]*intXYZ[2]
		sum += value
	}
	return sum
}

func main() {
	file, err := os.Open("2015/day02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	stringSlice, err := util.ReadToArray(file)

	if err != nil {
		log.Fatal(err)
	}

	resultPt1 := part1(stringSlice)
	resultPt2 := part2(stringSlice)

	fmt.Println(resultPt1)
	fmt.Println(resultPt2)
}
