package strings 

import (
	"log"
)


func Index(str string,matchstr string) int {

	var l1 ,l2 = len(str),len(matchstr) 
	var s,i = -1,0
	for {
		c := str[i]

		//reset
		if s >=0 && c != matchstr[i-s] {
			i = s+1
			s = -1
		}

		//init
		if s == -1 && c == matchstr[0] {
			s = i 
		}

		//break
		if s >= 0 && i - s == l2 - 1 || i > l1-1 {
			break
		}

		i++
	}

	return s
}


func Kmp(str string,matchstr string) int {

	var l1 ,l2 = len(str),len(matchstr) 

	log.Printf("intpu :%v %v",l1,l2)
	var s,i,ns = -1,0,-1
	for {
		c := str[i]

		//reset
		if s >=0 && c != matchstr[i-s] {
			if ns >= 0 {
				s = ns
				i = ns+1 
				ns = -1
				log.Printf("next %v %v %v",s,ns,i)
			} else {
				log.Printf("reset %v %v",s,ns)
				s,ns = -1,-1
			}
			
		}

		//init
		if s >= 0 &&  ns == -1 && c == matchstr[0] {
			ns = i
		}

		if s == -1 && c == matchstr[0] {
			s = i 
		}

		//break
		if s >= 0 && i - s == l2 - 1 {
			log.Printf("break %v %v %v",s,ns,i)
			break
		}

		
		if i >= l1-1 {
			//reset
			s = -1
			break
		} 

		i++
	}

	return s
}