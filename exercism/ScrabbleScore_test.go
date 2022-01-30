package exercism

import (
	"fmt"
	"testing"
)

func TestScrabbleScore(t *testing.T) {
	fmt.Println("Start TestScrabble")
	value1 := "cabbage"
	result1 := ScrabbleScore(value1)
	if result1 != 14 {
		t.Errorf("Not is expected value: %d", result1)
	}
}
