package queue

import (
	"testing"
	"log"
)


func TestPush(t *testing.T) {
	var q = NewQueue(2)

	q.Push(1)
	q.Push(2)
	log.Print(q)
	q.Push(3)

	log.Print(q)
}


func TestPop(t *testing.T) {
	var q = NewQueue(2)

	q.Push(1)
	// q.Push(2)

	log.Print(q)

	log.Printf("pop value %v",q.Pop())
	log.Print(q)
}

func TestCopy(t *testing.T ){
	var src []interface{} = []interface{}{1,2,3,4,5}
	log.Print(src)

	// var target = make([]interface{},10,10)
	// copy(src,0,5,target)
	// log.Print(target)

	var target = make([]interface{},10,10)
	copy(src,3,2,target)
	log.Print(target)
	
}


func TestPushPop(t *testing.T) {
	var q = NewQueue(4)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	log.Print(q)

	q.Pop()
	q.Pop()
	log.Print(q)

	q.Push(4)
	q.Push(5)
	q.Push(6)
	log.Print(q)
	q.Push(7)
	q.Push(8)
	q.Push(9)
	q.Push(10)
	q.Pop()
	q.Push(11)
	log.Print(q)
}