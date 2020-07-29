package main

import (
	"testing"
	"fmt"
)


func TestDeckErrorHandler(t *testing.T){
	fmt.Println("Check that error handler fires the message used to create it when an error occurs")

	message := "Testing deck Error Handler"
	var err error = deckError(message)
	result := err.Error()
	if result != message {
		t.Errorf("Expected %v but got %v", message, result)
	}
}


	func TestNewDeck(t *testing.T){
		fmt.Println("Testing the newDeck() function")
		fmt.Println("Test that newDeck returns a type of deck value")
		var i interface{} = newDeck()
		d, ok := i.(deck)
		if ok != true {
			t.Errorf("Expected newDeck to return a value of type deck")
		}

		fmt.Println("Test that the newDeck function returns a deck of 16 cards")
		expectedLength, actualLength := 16, len(d) 
		if expectedLength != actualLength {
			t.Errorf("Expected length of deck to be %v but got %v", expectedLength, actualLength)
		}

		fmt.Println("Test that the first card returned by the deck is an Ace of Spades")
		expectedCard, actualCard := "Ace of Spades", d[0]
		if expectedCard != actualCard {
			t.Errorf("Expected card type to be %v but got %v", expectedCard, actualCard)
		}

		fmt.Println("Test that the first card returned by the deck is a Three of Clubs")
		expectedCard, actualCard = "Three of Clubs", d[15]
		if expectedCard != actualCard {
			t.Errorf("Expected card type to be %v but got %v", expectedCard, actualCard)
		}
	}