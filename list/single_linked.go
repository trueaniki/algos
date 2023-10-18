package list

import (
	"fmt"
)

type SingleLinked struct {
	Data interface{}
	Next *SingleLinked
}

func (l *SingleLinked) Print() {
	fmt.Println(l.Data)
}

func (l *SingleLinked) Add(data interface{}) *SingleLinked {
	newL := &SingleLinked{
		Data: data,
		Next: nil,
	}
	l.Next = newL
	return newL
}

func (l *SingleLinked) AddMany(data ...interface{}) *SingleLinked {
	newL := l
	for _, value := range data {
		newL = newL.Add(value)
	}

	return newL
}

func NewSingleLinked(data interface{}) *SingleLinked {
	return &SingleLinked{
		Data: data,
		Next: nil,
	}
}

func (l *SingleLinked) PrintAll() {
	current := l
	for current.Next != nil {
		current.Print()
		current = current.Next
	}
	current.Print()
}

func (l *SingleLinked) At(index int) *SingleLinked {
	at := l
	for i := 0; i < index; i++ {
		at = at.Next
	}
	return at
}

func (l *SingleLinked) End() *SingleLinked {
	end := l
	for end.Next != nil {
		end = end.Next
	}
	return end
}

func (l *SingleLinked) Length() int {
	length := 0

	end := l
	for end != nil {
		length++
		end = end.Next
	}
	return length
}

func (l *SingleLinked) Reverse() *SingleLinked {
	end := l.End()
	current := l
	var prev *SingleLinked = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return end
}

func main() {
	start := NewSingleLinked(1)
	l := start

	l = l.Add(2)
	l = l.Add(3)

	l = l.AddMany(4, 5, 6)

	start.PrintAll()
	fmt.Printf("Length: %d\n", start.Length())

	reversed := start.Reverse()
	reversed.PrintAll()
}
