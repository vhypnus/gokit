package ring

import (
	// "log"
)

type Ring struct {

	//slice
	s []interface {}

	// head
	h int

	//capacity
	c int
}

func NewRing(capacity int) *Ring {
	var s = make([]interface{},2,capacity)
	var r = & Ring{s,0,capacity}
	return r
}

func (r *Ring) Push(v interface{}){
	r.s[r.h] = v

	r.h++
	if r.h == r.c {
		r.h = 0
	} 
}