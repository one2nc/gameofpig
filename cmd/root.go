package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "pig",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			cmd.Help()
			return
		}

		p1, err := strconv.Atoi(args[0])
		if err != nil || p1 < 1 || p1 > 100 {
			fmt.Println("Expect the strategy between 1 and 100")
			return
		}
		p1HoldScore := int8(p1)

		p2, err := strconv.Atoi(args[1])
		if err != nil || p2 < 1 || p2 > 100 {
			fmt.Println("Expect the strategy between 1 and 100")
			return
		}
		p2HoldScore := int8(p2)

		run(p1HoldScore, p2HoldScore)
	},
}

func init() {

	rootCmd.SetUsageFunc(nil)
	rootCmd.SetUsageTemplate(
		`pig - A command line tool to simulate a game of pig. It is a two-player game played with a 6-sided die.

Usage:
	pig [strategy] [strategy]

Args:
	strategy   The number between 1 to 100

Description:
	This command line application accepts two numbers between 1 to 100 as a positional argument. Strategies for player 1 and player 2, and performs the game of pig simulation on it. If no strategies are provided, then it will return an error. 

	If the number is out of range, the application will exit with an error message. Otherwise, it will perform the simulation and output the result.

Example usage:
	$ pig 10 15
	Result: Holding at   10 vs Holding at   15: wins: 3/10 (30.0%), losses: 7/10 (70.0%)
`,
	)

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

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

func run(p1HoldScore, p2HoldScore int8) {
	rand.Seed(time.Now().UnixNano())

	player1 := player{
		name:      "player1",
		holdScore: p1HoldScore,
	}

	player2 := player{
		name:      "player2",
		holdScore: p2HoldScore,
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
	fmt.Printf(
		"Result: Holding at %4d vs Holding at %4d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n",
		p1.holdScore,
		p2.holdScore,
		p1Wins,
		numberOfGames,
		p1WinsPer,
		numberOfGames-p1Wins,
		numberOfGames,
		100-p1WinsPer,
	)
}
