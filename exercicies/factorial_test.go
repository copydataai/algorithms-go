package exercicies

import (
	"fmt"
	"testing"
)

func TestFactorialIteracion(t *testing.T) {
	fmt.Println("Init Factorial Iteration")
	result1 := FactorialIteration(3)
	if result1 != 6 {
		t.Errorf("Not is value expected: %d", result1)
	}
}

func TestFactorialRecursive(t *testing.T) {
	fmt.Println("Init Factorial Recursive")
	result1 := FactorialRecusive(3)
	if result1 != 6 {
		t.Errorf("Not is value expected: %d", result1)
	}
}
