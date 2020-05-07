package stack 


type Stack struct {

	s []interface{}

	t int

	c int
}


func NewStack(capacity int) *Stack{
	var s = make([]interface{},capacity,capacity)
	var stack = &Stack{s,-1,capacity}
	return stack
}


func (s *Stack) Push(v interface{}) {

	if s.t == s.c -1 {
		var news = make([]interface{},2*s.c,2*s.c)
		copy(s.s,0,s.t,news)

		s.s = news
	}

	s.t++
	s.s[s.t] = v
}

func (s *Stack) Pop() interface{} {
	if s.t == -1 {
		//panic
	}
	var v = s.s[s.t]
	s.t --
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
