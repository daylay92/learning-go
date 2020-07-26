package main

import (
	"fmt"
)


func main(){
	d := newDeck()
	newD, remainderD := deal(d, 4)
	newD.print()
	remainderD.print()
	fmt.Println(newD.toString())
}