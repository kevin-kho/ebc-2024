package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pair struct {
	Enemy0 byte
	Enemy1 byte
}

type Triplet struct {
	Pair
	Enemy2 byte
}

func (p Pair) GetPotionCost() int {
	mp := map[byte]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
	}
	var extra int
	if p.Enemy0 != 'x' && p.Enemy1 != 'x' {
		extra = 2
	}

	return mp[p.Enemy0] + mp[p.Enemy1] + extra

}

func CreatePair(data [2]byte) Pair {
	enemy0 := data[0]
	enemy1 := data[1]

	return Pair{
		Enemy0: enemy0,
		Enemy1: enemy1,
	}
}

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

func SolvePartTwo(data []byte) int {
	var pairs []Pair
	for i := 1; i < len(data); i += 2 {
		a := data[i-1]
		b := data[i]

		p := CreatePair([2]byte{a, b})
		pairs = append(pairs, p)
	}

	var res int
	for _, p := range pairs {
		res += p.GetPotionCost()
	}

	return res

}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)

	res := SolvePartOne(dataPartOne)
	fmt.Println(res)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)
	res2 := SolvePartTwo(dataPartTwo)
	fmt.Println(res2)

}
