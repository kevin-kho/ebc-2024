package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func Solve(intArr []int) {
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
	Solve(intArr)
}

func SolvePartTwo(intArr []int) {
	Solve(intArr)
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

}
