package main

import (
	"fmt"
	"math/rand"
)

type Suite int
type Rank int

const (
	Heart Suite = iota
	Diamond
	Spade
	Club
)

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Rank Rank
	Suit Suite
}

type Deck []Card

func NewDeck() Deck {
	var deck Deck
	for r := range Ace + 1 {
		for s := range Club + 1 {
			var card Card = Card{Rank: r, Suit: s}
			deck = append(deck, card)
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d *Deck) Pop() Card {
	var card = (*d)[0]
	(*d) = (*d)[1:]
	return card
}

func (s Suite) String() string {
	return [...]string{"Heart", "Diamond", "Spade", "Club"}[s]
}

func (r Rank) String() string {
	return [...]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}[r]
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

func main() {
	var deck Deck = NewDeck()
	deck.Shuffle()
	fmt.Println("This is a stupid card game but its not about the game its about go")
	fmt.Println("So you pick the top card of the deck then the cpu picks a card and whoever has the highest card gets both")
	fmt.Println("Whoever has the most cards at the end of the deck wins really simple No Skill All Luck")
	var p1score int = 0
	var p2score int = 0
	for len(deck) > 0 {
		fmt.Printf("Player 1 Score: %d, Player 2 Score: %d\n", p1score, p2score)
		fmt.Print("Player 1 Press Enter to Pick A Card")
		fmt.Scanln()
		var top_card = deck.Pop()
		var second_card = deck.Pop()
		fmt.Printf("You picked a %s\n", top_card)
		fmt.Printf("CPU picked a %s\n", second_card)
		if top_card.Rank > second_card.Rank {
			fmt.Println("Your card is worth more you get both cards")
			p1score += 2
		} else if top_card.Rank < second_card.Rank {
			fmt.Println("Your card is worth less you lose your card to the CPU")
			p2score += 2
		} else {
			fmt.Println("Your cards are worth the same you both keep your card")
			p1score++
			p2score++
		}
	}
	if p1score > p2score {
		fmt.Println("YOU WIN")
	} else if p1score < p2score {
		fmt.Println("YOU LOSE")
	} else {
		fmt.Println("DRAW")
	}
}
