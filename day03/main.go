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
	Spots    [][]Spot
	Diggable map[models.Pos]bool
}

func CreateGrid(data []byte) Grid {
	var spots [][]Spot
	mp := make(map[models.Pos]bool)
	for y, row := range bytes.Split(data, []byte{'\n'}) {
		var rw []Spot
		for x, val := range row {
			var diggable bool
			if val == '#' {
				diggable = true
				mp[models.Pos{X: x, Y: y}] = true
			}

			rw = append(rw, Spot{
				Val:      0,
				Diggable: diggable,
			})

		}
		spots = append(spots, rw)

	}

	return Grid{
		Spots:    spots,
		Diggable: mp,
	}

}

func SolvePartOne(grid Grid) {

	deltas := models.GetDeltas()

	X := len(grid.Spots[0])
	Y := len(grid.Spots)

	for len(grid.Diggable) > 0 {
		var newSpots [][]Spot
		newDiggable := make(map[models.Pos]bool)
		for y, row := range grid.Spots {
			var spotRow []Spot
			for x, curr := range row {
				if !curr.Diggable {
					spotRow = append(spotRow, curr)
					continue
				}

				diggable := true
				newVal := curr.Val
				for _, d := range deltas {
					newX := x + d.X
					newY := y + d.Y
					// out of bounds
					if !(0 <= newX && newX < X) || !(0 <= newY && newY < Y) {
						continue
					}

					adj := grid.Spots[newY][newX]
					if common.IntAbs(adj.Val-curr.Val)+1 > 1 {
						diggable = false
					}
				}
				if diggable {
					newVal++
					newDiggable[models.Pos{X: x, Y: y}] = true

				}

				spotRow = append(spotRow, Spot{
					Val:      newVal,
					Diggable: diggable,
				})

				// grid.Spots[y][x] = Spot{
				// 	Val:      newVal,
				// 	Diggable: diggable,
				// }

			}
			newSpots = append(newSpots, spotRow)
		}
		grid.Spots = newSpots
		grid.Diggable = newDiggable
	}

	var count int
	for _, row := range grid.Spots {

		var vals []int
		for _, s := range row {
			if s.Val > 0 {
				count += s.Val
			}
			vals = append(vals, s.Val)
		}
		fmt.Println(vals)
	}
	fmt.Println(count)

}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	grid := CreateGrid(data)
	SolvePartOne(grid)

}
