package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	fmt.Println("Start TestQueue")
	var que Queue = *New(5)
	que.Enqueue(5)
	que.Enqueue(25)
	que.Enqueue(105)
	peek := que.Peek()
	if peek != 5 {
		t.Errorf("got %d expected", peek)
	}
}

func TestDequeue(t *testing.T) {
	fmt.Println("Start TestDequeue")
	var que Queue = *New(5)
	que.Enqueue(5)
	que.Enqueue(45)
	que.Enqueue(205)

	deque := que.Dequeue()
	if deque != 5 {
		t.Errorf("got %d expected", deque)
	}
	peek := que.Peek()
	if peek != 45 {
		t.Errorf("expected peek with value: 25 but is %d", peek)
	}
}
