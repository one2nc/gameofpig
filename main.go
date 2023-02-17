package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Story 1
// Player 1 uses a strategy of always holding after accumulating a score of at least 10,
// while Player 2 uses a strategy of always holding after reaching a sum of at least 15.

type player struct {
	name      string
	holdScore int8
	score     int8
	turnScore int8
	turnRolls []int8
}

var winningScore = int8(100)
var numberOfGames = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	player1 := player{
		name:      "player1",
		holdScore: 2,
	}

	player2 := player{
		name:      "player2",
		holdScore: 50,
	}

	wins := make(map[string]int)
	for i := 0; i < numberOfGames; i++ {
		// reset the score
		player1.score = 0
		player2.score = 0

		p := &player1
		for {
			p.score += p.playTurn()

			if p.score >= winningScore {
				break
			}

			// switch the turn
			if p.name == "player1" {
				p = &player2
			} else {
				p = &player1
			}
		}
		wins[p.name]++
	}

	printRatio(player1, player2, wins)
}

func rollDice() int8 {
	value := int8(rand.Intn(6) + 1)
	return value
}

func (p player) getDecision() string {
	if p.score+p.turnScore >= winningScore {
		return "hold"
	} else if p.turnScore >= p.holdScore {
		return "hold"
	} else {
		return "roll"
	}
}

func (p *player) playTurn() int8 {
	p.turnScore = 0
	p.turnRolls = []int8{}
	for p.getDecision() == "roll" {
		diceValue := rollDice()
		p.turnRolls = append(p.turnRolls, diceValue)
		if diceValue == 1 {
			return 0
		}
		p.turnScore += diceValue
	}
	return p.turnScore
}

func printRatio(p1, p2 player, wins map[string]int) {
	p1Wins := wins[p1.name]
	p1WinsPer := 100 * float64(p1Wins) / float64(numberOfGames)
	fmt.Printf("Holding at %4d vs Holding at %4d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n",
		p1.holdScore, p2.holdScore, p1Wins, numberOfGames, p1WinsPer, numberOfGames-p1Wins, numberOfGames, 100-p1WinsPer)
}
