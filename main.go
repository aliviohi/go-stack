package main

import (
	"fmt"
	"time"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value, true
}

func (s Stack) Print() {
	fmt.Println("Stack (bottom -> top):", s.items)
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stack := Stack{}

	fmt.Println("Input slice:", input)

	fmt.Println("\nPushing 10 items into stack...")
	for _, v := range input {
		stack.Push(v)
		fmt.Printf("Pushed: %d\n", v)
		stack.Print()
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nPopping items from stack...")
	for {
		popped, ok := stack.Pop()
		if !ok {
			fmt.Println("Stack is empty.")
			break
		}

		fmt.Println("Popped value:", popped)
		stack.Print()
		time.Sleep(1 * time.Second)
	}
}
