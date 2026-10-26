package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolveMin(intArr []int) {
	mn := intArr[0]
	for _, val := range intArr {
		mn = min(mn, val)
	}

	var res int
	for _, val := range intArr {
		res += (val - mn)
	}

	fmt.Println(res)

}

func SolvePartOne(intArr []int) {
	SolveMin(intArr)
}

func SolvePartTwo(intArr []int) {
	SolveMin(intArr)
}

func SolvePartThree(intArr []int) {
	// Median will minimize the absolute devince
	slices.Sort(intArr)
	l := len(intArr)/2 - 1
	r := len(intArr) / 2
	med := (intArr[l] + intArr[r]) / 2

	var res int
	for _, val := range intArr {
		res += common.IntAbs(val - med)
	}

	fmt.Println(res)
}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)
	intArrPartOne, err := common.ParseIntArray(dataPartOne, []byte{'\n'})
	if err != nil {
		log.Fatal(err)
	}
	SolvePartOne(intArrPartOne)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)
	intArrPartTwo, err := common.ParseIntArray(dataPartTwo, []byte{'\n'})
	if err != nil {
		log.Fatal(err)
	}
	SolvePartTwo(intArrPartTwo)

	dataPartThree, err := common.ReadInput("inputPartThree.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartThree = common.TrimNewLineSuffix(dataPartThree)
	intArrPartThree, err := common.ParseIntArray(dataPartThree, []byte{'\n'})
	if err != nil {
		log.Fatal(err)
	}
	SolvePartThree(intArrPartThree)

}
