// 04_basic_data_structures.go
// Demonstrates basic data structures in Go: stack, word frequency map, and Person struct.

package main

import (
	"fmt"
	"strings"
)

// Integer stack implementation
// ---------------------------
type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

func (s *Stack) Print() {
	fmt.Println("Stack:", s.data)
}

// Word frequency map
// ------------------
func WordFreq(s string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		freq[w]++
	}
	return freq
}

// Person struct and finding oldest
// -------------------------------
type Person struct {
	Name string
	Age  int
}

func PrintOldest(people []Person) {
	if len(people) == 0 {
		fmt.Println("No people provided.")
		return
	}
	oldest := people[0]
	for _, p := range people[1:] {
		if p.Age > oldest.Age {
			oldest = p
		}
	}
	fmt.Printf("Oldest: %s (%d years old)\n", oldest.Name, oldest.Age)
}

func main() {
	// Stack demo
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Print()
	val, ok := s.Pop()
	if ok {
		fmt.Println("Popped:", val)
	}
	s.Print()

	// Word frequency demo
	text := "go go is fun and go is easy"
	freq := WordFreq(text)
	fmt.Println("Word frequencies:", freq)

	// Person demo
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Carol", Age: 35},
	}
	PrintOldest(people)
}
