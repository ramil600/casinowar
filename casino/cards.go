package casino

import (
	"math/rand"
	"time"
)

type Suits int
type Ranks int

const (
	Clubs Suits = iota
	Diamonds
	Hearts
	Spades
)
const (
	Two Ranks = 2 + iota
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

const NumDecks = 6

// mranks and msuits implement readability
// fmt.Print( mranks[myDeck[i].Rank], " of ", msuits[myDeck[i].Suit], "\n" )
var mranks = map[Ranks]string{Two: "Two", Three: "Three", Four: "Four", Five: "Five", Six: "Six", Seven: "Seven",
	Eight: "Eight", Nine: "Nine", Ten: "Ten", Jack: "Jack", Queen: "Queen", King: "King", Ace: "Ace"}

var msuits = map[Suits]string{Clubs: "Clubs", Diamonds: "Diamonds", Hearts: "Hearts", Spades: "Spades"}

type Card struct {
	Suit Suits
	Rank Ranks
}

//Deck of Cards
type Deck []Card

//NewDeck factory for a deck of cards
func NewDeck() *Deck {
	deck := make(Deck, 52*NumDecks)
	for h := 0; h < NumDecks; h++ {
		for i := Clubs; i <= Spades; i++ {
			for j := Two; j <= Ace; j++ {
				deck[(int(i)*13+int(j-2))+(52*h)] = Card{
					Suit: i,
					Rank: j,
				}
			}
		}
	}
	return &deck
}

//Shuffle cards, otherwise sorted by rank, suit
func (d *Deck) Shuffle() {
	n := 312 // Times to shuffle two random cards
	for i := 0; i < n; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		j := r.Intn(52 * NumDecks)

		k := rand.Intn(52 * NumDecks)

		if j != k {
			(*d)[j].Rank, (*d)[k].Rank = (*d)[k].Rank, (*d)[j].Rank
			(*d)[j].Suit, (*d)[k].Suit = (*d)[k].Suit, (*d)[j].Suit
		}

	}

}
