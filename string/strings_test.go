package strings


import (
	"testing"
	"log"
)


func TestIndex(t *testing.T ){
	var str = "helloword"

	var matchstr = "ow"

	log.Print(Index(str,matchstr))
}

func TestKmp(t *testing.T) {
	var str = "helloword"

	var matchstr = "ow"

	log.Print(Kmp(str,matchstr))

	matchstr = "owr"

	log.Print(Kmp(str,matchstr))

	matchstr = "elr"

	log.Print(Kmp(str,matchstr))

}