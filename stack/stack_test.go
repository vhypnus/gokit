package stack

import (
	"testing"
	"log"
)


func TestPush(t *testing.T) {


	var s = NewStack(4)
	s.Push(1)
	s.Push(2)
	log.Print(s)

	s.Pop()
	log.Print(s)

	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	log.Print(s)
	var v int = s.Pop().(int)
	log.Print(v)
}