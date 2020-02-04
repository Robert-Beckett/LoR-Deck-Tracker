package main

import (
	"encoding/json"
	"fmt"
)

type board struct {
	PlayerName   string
	OpponentName string
	GameState    string
	Screen       screen
	Rectangles   []fieldCard
}

type screen struct {
	ScreenWidth  int
	ScreenHeight int
}

type fieldCard struct {
	LocalPlayer bool
	TopLeftX    int
	TopLeftY    int
	Width       int
	Height      int
	Card        card
}

var playerDeck deck

// Store the discard as an array to track order.
// The same will go for the field, but they'll contain
// maps that contain cards, to track board position, etc.
var playerDiscard []card
var player []int

var opponent []int

func getBoard() (board, error) {
	data, err := queryClient("positional-rectangles")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	var fieldCards board
	err = json.Unmarshal([]byte(data), &fieldCards)
	if err != nil {
		panic(err)
	}
	fmt.Println(fieldCards)

	return fieldCards, err
}

func initDeck(deck deck) {
	playerDeck = deck
}

// Will return an array of all cards that were not
// on the board or handalready, and then an array of all
// cards that were on the board or hand, but have since been
// removed.
func boardChanges(oldBoard, newBoard board) ([]card, []card) {
	var newCards []card
	var removedCards []card

	return newCards, removedCards
}

func handleRemovedCard(removedCard card) {

}

func handleAddedCard(newCard fieldCard) {
	if newCard.LocalPlayer {
		playerDeck.CardsInDeck[newCard.Card.CardCode]--
		fmt.Println("You played:", newCard.Card.Name)
	}
}