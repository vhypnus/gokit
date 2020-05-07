package heap

import (
	"log"
)

type Heap struct {

	s []interface{}

	// ""

	l int

	c int

	// type max(0) or min() > 0)
	t byte
}


func NewMaxHeap(capacity int) *Heap {
	var s = make([]interface{},capacity,capacity)
	var h = &Heap{s,0,capacity,0}
	return h
}

func NewMinHeap(capacity int) *Heap {
	var s = make([]interface{},capacity,capacity)
	var h = &Heap{s,0,capacity,1}
	return h
}


func (h *Heap) Push(v interface{}) {

	if h.l == h.c {
		var newhs = make([]interface{},2*h.c,2*h.c)
		copy(h.s,0,h.l-1,newhs)
		h.s = newhs
		h.l = h.l
		h.c = cap(newhs)
	}

	h.s[h.l] = v
	h.l++


	maxheapify(h.s,h.l-1)
}

func (h *Heap) Pop(v interface{}) interface{} {

	return nil
}


func maxheapify(s []interface{},i int) {

	if i == 0 {
		return 
	}
	
	var p int
	if i & 1 == 0 {
		p = (i-2)/2
	} else {
		p = (i - 1)/2
	}

	var v_1,v_2 int = s[i].(int),s[p].(int)
	if  v_1 > v_2 {
		temp := s[p]
		s[p] = s[i]
		s[i] = temp
	}

	log.Print(s)
	maxheapify(s,p)

}


func copy(src [] interface{},s ,e int,target []interface{}){

	var l,j = len(src),0
	
	for i := 0; i < l; i ++ {
		if e < s {
			if i > l - s-1 {
				j = i - (l - s)
			} else {
				j = s +i
			}
		} else {
			j = i
		}

		target[i] = src[j]
	}
}
