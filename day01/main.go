package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Group struct {
	Enemies []byte
}

func (t Group) GetPotionCost() int {
	mp := map[byte]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
	}
	var extra int
	switch len(t.Enemies) {
	case 2:
		extra = 2
	case 3:
		extra = 6
	}

	var res int
	for _, e := range t.Enemies {
		res += mp[e]
	}
	res += extra

	return res

}

func CreateGroup(data []byte) Group {

	var enemies []byte
	for _, d := range data {
		if d == 'x' {
			continue
		}
		enemies = append(enemies, d)
	}

	return Group{
		Enemies: enemies,
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
	var groups []Group
	for i := 1; i < len(data); i += 2 {
		a := data[i-1]
		b := data[i]

		g := CreateGroup([]byte{a, b})
		groups = append(groups, g)
	}

	var res int
	for _, g := range groups {
		res += g.GetPotionCost()
	}

	return res

}

func SolvePartThree(data []byte) int {
	var groups []Group
	for i := 2; i < len(data); i += 3 {
		a := data[i-2]
		b := data[i-1]
		c := data[i]

		t := CreateGroup([]byte{a, b, c})
		groups = append(groups, t)
	}

	var res int
	for _, g := range groups {
		res += g.GetPotionCost()
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

	dataPartThree, err := common.ReadInput("inputPartThree.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartThree = common.TrimNewLineSuffix(dataPartThree)
	res3 := SolvePartThree(dataPartThree)
	fmt.Println(res3)

}
