package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(data []byte) int {
	var res int
	mp := map[byte]int{
		'A': 0,
		'B': 1,
		'C': 3,
	}

	for _, b := range data {
		res += mp[b]
	}

	return res

}

func main() {
	data, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	res := SolvePartOne(data)
	fmt.Println(res)

}
