package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Spot struct {
	Val      int
	Diggable bool
}

type Grid struct {
	Spots    map[models.Pos]Spot
	Diggable map[models.Pos]bool
}

func CreateGrid(data []byte) Grid {
	spots := make(map[models.Pos]Spot)
	mp := make(map[models.Pos]bool)
	for y, row := range bytes.Split(data, []byte{'\n'}) {
		for x, val := range row {
			var diggable bool
			if val == '#' {
				diggable = true
				mp[models.Pos{X: x, Y: y}] = true
			}

			spots[models.Pos{X: x, Y: y}] = Spot{Val: 0, Diggable: diggable}

		}

	}

	return Grid{
		Spots:    spots,
		Diggable: mp,
	}

}

func Solve(grid Grid, deltas []models.Pos) {

	var X int
	var Y int
	for p := range grid.Spots {
		X = max(X, p.X)
		Y = max(Y, p.Y)
	}
	X++
	Y++

	for len(grid.Diggable) > 0 {
		newSpots := make(map[models.Pos]Spot)
		newDiggable := make(map[models.Pos]bool)
		for p, s := range grid.Spots {
			if !s.Diggable {
				newSpots[p] = s
				continue
			}

			diggable := true
			newVal := s.Val
			for _, d := range deltas {
				newX := p.X + d.X
				newY := p.Y + d.Y
				// out of bounds
				// if !(0 <= newX && newX < X) || !(0 <= newY && newY < Y) {
				// 	continue
				// }

				adj := grid.Spots[models.Pos{X: newX, Y: newY}]
				if common.IntAbs(adj.Val-s.Val)+1 > 1 {
					diggable = false
				}
			}
			if diggable {
				newVal++
				newDiggable[models.Pos{X: p.X, Y: p.Y}] = true
			}

			newSpots[models.Pos{X: p.X, Y: p.Y}] = Spot{
				Val:      newVal,
				Diggable: diggable,
			}

		}

		grid.Spots = newSpots
		grid.Diggable = newDiggable
	}

	var count int
	for _, s := range grid.Spots {
		if s.Val > 0 {
			count += s.Val
		}
	}
	fmt.Println(count)

}

func SolvePartOne(grid Grid) {
	Solve(grid, models.GetDeltas())
}

func SolvePartTwo(grid Grid) {
	Solve(grid, models.GetDeltas())
}

func SolvePartThree(grid Grid) {
	var deltas []models.Pos
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			deltas = append(deltas, models.Pos{X: x, Y: y})
		}
	}

	Solve(grid, deltas)

}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	gridPartOne := CreateGrid(data)
	SolvePartOne(gridPartOne)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)
	gridPartTwo := CreateGrid(dataPartTwo)
	SolvePartTwo(gridPartTwo)

	dataPartThree, err := common.ReadInput("inputPartThree.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartThree = common.TrimNewLineSuffix(dataPartThree)
	gridPartThree := CreateGrid(dataPartThree)
	SolvePartThree(gridPartThree)

}
