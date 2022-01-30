package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	fmt.Println("Start TestStack")
	stack := New(5)
	stack.Add(5)
	stack.Add(25)
	stack.Add(55)
	peek := stack.Peek()
	if peek != 55 {
		t.Errorf("Not is equal to last %v", peek)
	}
}

func TestRemoveStack(t *testing.T) {
	fmt.Println("Stack TestRemoveStack")
	stack := New(10)
	stack.Add(5)
	stack.Add(50)
	stack.Add(500)
	remove := stack.Remove()
	if remove != 500 {
		t.Errorf("Not are remove %d", remove)
	}
	peek := stack.Peek()
	if peek != 50 {
		t.Errorf("Not are remove %d", peek)
	}
}
