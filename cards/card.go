package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"math/rand"
	"time"
)

type deck []string
type deckError string

func (d deckError) Error() string {
	return string(d)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func createDirIfNotExists(dirname string) (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dirnameAbs := filepath.Join(root, dirname)
	if _, err := os.Stat(dirnameAbs); err != nil {
		if os.IsNotExist(err) {
			return dirnameAbs, os.Mkdir(dirnameAbs, 0666)
		}
	}
	return dirnameAbs, nil
}

func (d deck) saveToFile(filename string) error {
	message := []byte(d.toString())
	dirname, err := createDirIfNotExists("/cards.chest")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(dirname, filename+".txt"), message, 0666)
	return err
}

func readDeckFromFile(filename string) (de deck, e error) {
	root, err := os.Getwd()
	if err != nil {
		e = err
		return
	}
	bs, err := ioutil.ReadFile(filepath.Join(root, "/cards.chest", filename+".txt"))
	if err != nil {
		e = deckError("Error fetching deck from file, the file named " + filename + " probably doesn't exist, check that cards.chest directory already exists")
		return
	}
	de = deck(strings.Split(string(bs), ","))
	return
}

func (d deck) stuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newI := r.Intn(len(d) - 1)
		d[i] = d[newI]
	}
}