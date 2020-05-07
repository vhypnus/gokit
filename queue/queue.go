package queue

import (
	// "log"
)

type Queue struct {

	s []interface{}

	h int

	t int

	l int

	c int
}


func NewQueue(capacity int) *Queue{
	var s = make([]interface{},capacity,capacity)

	var q =  &Queue{s,-1,-1,0,capacity}
	return q 
}


func (q *Queue) Push(v interface{}) {

	//resize
	if q.l == q.c {

		var newqs = make([]interface{},2*q.c,2*q.c)
		copy(q.s,q.h,q.t,newqs)
		q.s = newqs
		q.h = 0
		q.t = q.l-1
		q.l = q.l
		q.c = cap(newqs)

	}

	q.t++
	if q.t == q.c {
		q.t = 0
	} 
	q.s[q.t] = v
	q.l++
	
	//init head
	if q.h == -1 {
		q.h = 0
	}

}

func (q *Queue) Pop() interface{} {

	var v interface{}
	if q.l == 0 {
		//panic
	}else {
		q.h++
		v = q.s[q.h]
		q.l--
		if q.h == q.c{
			q.h = 0
		} 
	}

	return v

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


