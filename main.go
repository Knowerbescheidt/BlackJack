package main

import (
	"fmt"
	"strings"

	deck "github.com/Knowerbescheidt/Deck-of-cards"
)

//Reimplementing 7:10

type Hand []deck.Card

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **Hidden**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Shuffle(gs Gamestate) Gamestate {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gs Gamestate) Gamestate {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 7)
	ret.Dealer = make(Hand, 0, 7)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.Player = append(ret.Player, card)
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

func main() {

	var gs Gamestate
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("Player:", gs.Player)
		fmt.Println("Dealer:", gs.Dealer.DealerString())
		fmt.Println("What will you do?(h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}

	}

	for dealer.Score() < 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("********Final Score**********")
	fmt.Println("Player hand: ", player.String(), "\n Player Score:", pScore)
	fmt.Println("Dealer hand: ", dealer.String(), "\n Dealer Score:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted, and loose, Looser")
	case dScore > 21:
		fmt.Println("Dealer busted, and loose, Looser")
	case pScore > dScore:
		fmt.Println("You win Congrats")
	case dScore > pScore:
		fmt.Println("You loose...Try Again!")
	case dScore == pScore:
		fmt.Println("Draw")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type Gamestate struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *Gamestate) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("A valid Gamestate can not be found")
	}
}

func clone(gs Gamestate) Gamestate {
	ret := Gamestate{
		Deck:   make([]deck.Card, len(gs.Deck)),
		Turn:   gs.Turn,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	//hard copy not mix values or change old game states (Inplace bei pandas)
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)

	return ret
}
