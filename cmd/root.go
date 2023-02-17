package cmd

import (
	"fmt"
	"gameofpig/domain"
	"os"
	"strconv"

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

var winningScore = int8(100)
var numberOfGames = 10

// Story 1
// Player 1 uses a strategy of always holding after accumulating a score of at least 10,
// while Player 2 uses a strategy of always holding after reaching a sum of at least 15.
func run(p1HoldScore, p2HoldScore int8) {
	player1 := domain.Player{
		Name:      "player1",
		HoldScore: p1HoldScore,
	}

	player2 := domain.Player{
		Name:      "player2",
		HoldScore: p2HoldScore,
	}

	dice := domain.Dice{}

	wins := make(map[string]int)
	for i := 0; i < numberOfGames; i++ {
		// reset the score
		player1.Score = 0
		player2.Score = 0

		p := &player1
		for {
			p.Score += p.PlayTurn(dice, winningScore)

			if p.Score >= winningScore {
				break
			}

			// switch the turn
			if p.Name == "player1" {
				p = &player2
			} else {
				p = &player1
			}
		}
		wins[p.Name]++
	}

	printRatio(player1, player2, wins)
}

func printRatio(p1, p2 domain.Player, wins map[string]int) {
	p1Wins := wins[p1.Name]
	p1WinsPer := 100 * float64(p1Wins) / float64(numberOfGames)
	fmt.Printf(
		"Result: Holding at %4d vs Holding at %4d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n",
		p1.HoldScore,
		p2.HoldScore,
		p1Wins,
		numberOfGames,
		p1WinsPer,
		numberOfGames-p1Wins,
		numberOfGames,
		100-p1WinsPer,
	)
}
