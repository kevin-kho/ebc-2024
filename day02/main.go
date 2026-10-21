package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
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

	inscription := strArr[len(strArr)-1]

	return Input{
		Words:       wordArr,
		Inscription: inscription,
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

func main() {
	dataPartOne, err := common.ReadInput("inputPartOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	dataPartOne = common.TrimNewLineSuffix(dataPartOne)
	in := ParseInput(dataPartOne)
	res := SolvePartOne(in)
	fmt.Println(res)

}
