package heap

import(
	"testing"
	"log"
)


func TestPush(t *testing.T) {
	var h = NewMaxHeap(2)

	h.Push(1)
	h.Push(2)

	// log.Print(h)

	h.Push(3)
	h.Push(4)
	// log.Print(h)

	h.Push(5)
	log.Print(h)
}