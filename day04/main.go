package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(intArr []int) {
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

func main() {
	data, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	intArr, err := common.ParseIntArray(data, []byte{'\n'})
	if err != nil {
		log.Fatal(err)
	}

	SolvePartOne(intArr)

}
