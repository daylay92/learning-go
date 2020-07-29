package main

import (
	"fmt"
)


func main(){
	d := newDeck()
	newD, remainderD := deal(d, 4)
	err:= remainderD.saveToFile("remainder")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newD.toString())
	de, err := readDeckFromFile("remainder")
	if err != nil {
		fmt.Println(err)
	}
	de.stuffle()
	de.print()
}