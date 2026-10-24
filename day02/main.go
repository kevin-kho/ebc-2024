package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Input struct {
	Words       []string
	Inscription string
}

func ParseInput(data []byte) Input {
	strArr := strings.Split(string(data), "\n")
	words := strArr[0]
	words = strings.TrimPrefix(words, "WORDS:")
	wordArr := strings.Split(words, ",")

	inscriptionArr := strArr[2:]
	inscription := strings.Join(inscriptionArr, " ")

	return Input{
		Words:       wordArr,
		Inscription: inscription,
	}
}

type InputGrid struct {
	Words []string
	Grid  [][]byte
}

func ParseInputGrid(data []byte) InputGrid {
	strArr := strings.Split(string(data), "\n")
	words := strArr[0]
	words = strings.TrimPrefix(words, "WORDS:")
	wordArr := strings.Split(words, ",")

	inscriptionArr := strArr[2:]
	var grid [][]byte
	for _, s := range inscriptionArr {
		grid = append(grid, []byte(s))

	}

	return InputGrid{
		Words: wordArr,
		Grid:  grid,
	}
}

func SolvePartOne(in Input) int {
	mp := make(map[int][]string)
	for _, w := range in.Words {
		mp[len(w)] = append(mp[len(w)], w)
	}

	var count int
	for ln, words := range mp {
		// Build a set of words for the current length
		set := make(map[string]bool)
		for _, w := range words {
			set[w] = true
		}

		// double pointer the inscription
		l := 0
		r := ln
		for r < len(in.Inscription) {
			word := in.Inscription[l:r]
			if set[word] {
				count++
			}

			// move double pointers
			l++
			r++
		}

	}

	return count

}

func SolvePartTwo(in Input) int {
	mp := make(map[int][]string)
	for _, w := range in.Words {
		mp[len(w)] = append(mp[len(w)], w)

		rev := []byte(w)
		slices.Reverse(rev)
		mp[len(w)] = append(mp[len(w)], string(rev))
	}

	indexes := make(map[int]bool)
	for ln, words := range mp {
		// Build a set of words for the current length
		set := make(map[string]bool)
		for _, w := range words {
			set[w] = true
		}

		// double pointer the inscription
		l := 0
		r := ln
		for r < len(in.Inscription) {
			word := in.Inscription[l:r]
			if set[word] {
				for i := l; i < r; i++ {
					indexes[i] = true
				}
			}

			// move double pointers
			l++
			r++
		}

	}

	return len(indexes)

}

func SolvePartThree(in InputGrid) {

	// Build Word Map
	mp := make(map[int][]string)
	for _, w := range in.Words {
		mp[len(w)] = append(mp[len(w)], w)

		rev := []byte(w)
		slices.Reverse(rev)
		mp[len(w)] = append(mp[len(w)], string(rev))
	}

	seen := make(map[models.Pos]bool)
	// Assess Rows
	for y, row := range in.Grid {
		for ln, words := range mp {
			set := make(map[string]bool)
			for _, w := range words {
				set[w] = true
			}
			l := 0
			r := ln
			for r < len(row) {
				word := string(row[l:r])
				if set[word] {
					for x := l; x < r; x++ {
						seen[models.Pos{X: x, Y: y}] = true
					}
				}
				l++
				r++
			}
		}
	}

	// Assess Cols
	cols := make(map[int][]byte) // key: x, val: all y
	X := len(in.Grid[0])
	Y := len(in.Grid)
	for x := range X {
		var values []byte
		for y := range Y {
			val := in.Grid[y][x]
			values = append(values, val)
		}
		cols[x] = values
	}

	for x, row := range cols {
		for ln, words := range mp {
			set := make(map[string]bool)
			for _, w := range words {
				set[w] = true
			}
			l := 0
			r := ln
			for r < len(row) {
				word := string(row[l:r])
				if set[word] {
					for y := l; y < r; y++ {
						seen[models.Pos{X: x, Y: y}] = true
					}
				}
				l++
				r++
			}
		}
	}

	fmt.Println(len(seen))

}

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)
	inPartOne := ParseInput(dataPartOne)
	res := SolvePartOne(inPartOne)
	fmt.Println(res)

	dataPartTwo, err := common.ReadInput("inputPartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartTwo = common.TrimNewLineSuffix(dataPartTwo)
	inPartTwo := ParseInput(dataPartTwo)
	res2 := SolvePartTwo(inPartTwo)
	fmt.Println(res2)

	dataPartThree, err := common.ReadInput("inputPartThree.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartThree = common.TrimNewLineSuffix(dataPartThree)
	inPartThree := ParseInputGrid(dataPartThree)
	SolvePartThree(inPartThree)

}
