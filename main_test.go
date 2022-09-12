package main

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	type Condition struct {
		Input  int
		Result int
	}
	var stamps = [...]ScoreStamp{
		{Offset: 0, Score: Score{Home: 0, Away: 0}},
		{Offset: 2, Score: Score{Home: 1, Away: 0}},
		{Offset: 3, Score: Score{Home: 2, Away: 0}},
		{Offset: 6, Score: Score{Home: 2, Away: 1}},
		{Offset: 8, Score: Score{Home: 2, Away: 2}},
	}
	var offsets = []Condition{
		{Input: -3, Result: 0},
		{Input: 0, Result: 0},
		{Input: 1, Result: 0},
		{Input: 2, Result: 1},
		{Input: 4, Result: 2},
		{Input: 5, Result: 2},
		{Input: 6, Result: 3},
		{Input: 8, Result: 4},
		{Input: 10, Result: 4},
	}

	fmt.Printf("Testing getScore function with ScoreStamp array %v\n", stamps)

	for i, offset := range offsets {
		t.Run(fmt.Sprintf("Test number: %v\n", i), func(t *testing.T) {
			res := getScore(stamps[:], offset.Input)
			if res != stamps[offset.Result].Score {
				t.Error(
					"For offset input", offset.Input,
					"expected score", stamps[offset.Result].Score,
					"got score", res,
				)
			}
		})
	}
}
