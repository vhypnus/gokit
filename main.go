package main 

import(
	"log"
)

func main(){

	var i interface{} = 1

	var j interface{} = 1

	log.Print(i == j)
}