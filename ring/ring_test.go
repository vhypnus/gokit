package ring

import (
	"testing"
	"log"
)


func TestPush(t *testing.T) {

	var r = NewRing(2)

	r.Push(1)
	r.Push(2)

	log.Print(r)

	r.Push(3)
	log.Print(r)
}