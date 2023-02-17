package cmd

import (
	"fmt"
	"gameofpig/domain"
	"os"
	"strconv"
	"strings"

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

		p1, p2 := args[0], args[1]
		if strings.Contains(p1, "-") && strings.Contains(p2, "-") {
			playMultiStrategyWithMultiStrategy(p1, p2)
			return
		}

		p1HoldScore, err := strconv.Atoi(p1)
		if err != nil || p1HoldScore < 1 || p1HoldScore > 100 {
			fmt.Println("expect the strategy between 1 and 100")
			return
		}

		if strings.Contains(p2, "-") {
			playSingleStrategyWithMultiStrategy(p1HoldScore, p2)
			return
		}

		p2HoldScore, err := strconv.Atoi(p2)
		if err != nil || p2HoldScore < 1 || p2HoldScore > 100 {
			fmt.Println("expect the strategy between 1 and 100")
			return
		}

		p1Wins, gamesPlayed := play(p1HoldScore, p2HoldScore)
		printRatio1(p1HoldScore, p2HoldScore, p1Wins, gamesPlayed)
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
func play(p1HoldScore, p2HoldScore int) (int, int) {
	player1 := domain.Player{
		Name:      "player1",
		HoldScore: int8(p1HoldScore),
	}

	player2 := domain.Player{
		Name:      "player2",
		HoldScore: int8(p2HoldScore),
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

	return wins[player1.Name], numberOfGames
}

// Story 2
// Player 1 has a fixed strategy of always holding after accumulating a score of at least 21.
// Player 2 changes their strategy for each set of games between 1 to 100:
func playSingleStrategyWithMultiStrategy(p1HoldScore int, p2HoldScore string) {
	p2Start, p2End, err := extractRange(p2HoldScore)
	if err != nil {
		fmt.Println(err)
		return
	}

	for p2HoldScore := p2Start; p2HoldScore <= p2End; p2HoldScore++ {
		if p2HoldScore == p1HoldScore {
			continue
		}

		p1Wins, gamesPlayed := play(p1HoldScore, p2HoldScore)
		printRatio1(p1HoldScore, p2HoldScore, p1Wins, gamesPlayed)
	}
}

// Story 3
// This story is an extension of previous story, where both players change their strategies. In story 2, Player 1 had fixed strategy and Player 2 used different strategies. In this story, let's allow both players to change their strategies.

// Thus, there will be 100 strategies for Player 1 (from "hold until 1" to "hold until 100") and 99 strategies for Player 2. Each of these strategies will be played against each other in a match. Each such match will have 10 games. Thus, there will be 99,000 games (100 strategies * 99 strategies * 10 games per match) in total.

// Now, instead of printing 9900 output lines (as per Story 2's logic), let's do something different. Let's calculate the probability of win rates for strategies and print that.
func playMultiStrategyWithMultiStrategy(p1HoldScore, p2HoldScore string) {
	p1Start, p1End, err := extractRange(p1HoldScore)
	if err != nil {
		fmt.Println(err)
		return
	}

	p2Start, p2End, err := extractRange(p2HoldScore)
	if err != nil {
		fmt.Println(err)
		return
	}

	for p1HoldScore := p1Start; p1HoldScore <= p1End; p1HoldScore++ {
		p1WinsTotal := 0
		gamesPlayedTotal := 0
		for p2HoldScore := p2Start; p2HoldScore <= p2End; p2HoldScore++ {
			if p1HoldScore == p2HoldScore {
				continue
			}

			p1Wins, gamesPlayed := play(p1HoldScore, p2HoldScore)
			p1WinsTotal += p1Wins
			gamesPlayedTotal += gamesPlayed
		}
		printRatio2(p1HoldScore, p1WinsTotal, gamesPlayedTotal)
	}
}

func printRatio1(p1HoldScore, p2HoldScore, p1Wins, numberOfGames int) {
	p1WinsPer := 100 * float64(p1Wins) / float64(numberOfGames)
	fmt.Printf(
		"Result: Holding at %4d vs Holding at %4d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n",
		p1HoldScore,
		p2HoldScore,
		p1Wins,
		numberOfGames,
		p1WinsPer,
		numberOfGames-p1Wins,
		numberOfGames,
		100-p1WinsPer,
	)
}

func printRatio2(p1HoldScore, p1Wins, numberOfGames int) {
	p1WinsPer := 100 * float64(p1Wins) / float64(numberOfGames)
	fmt.Printf(
		"Result: Wins, losses staying at k = %4d: %d/%d (%0.1f%%), %d/%d (%0.1f%%)\n",
		p1HoldScore,
		p1Wins,
		numberOfGames,
		p1WinsPer,
		numberOfGames-p1Wins,
		numberOfGames,
		100-p1WinsPer,
	)
}

func extractRange(strategyRange string) (int, int, error) {
	ranngeArr := strings.Split(strategyRange, "-")
	if len(ranngeArr) > 2 {
		return 0, 0, fmt.Errorf("expect the strategy in range format e.g. 1-100")
	}

	start, err := strconv.Atoi(ranngeArr[0])
	if err != nil || start < 1 || start > 100 {
		return 0, 0, fmt.Errorf("expect the strategy between 1 and 100")
	}

	end, err := strconv.Atoi(ranngeArr[1])
	if err != nil || end < 1 || end > 100 {
		return 0, 0, fmt.Errorf("expect the strategy between 1 and 100")
	}

	if start >= end {
		return 0, 0, fmt.Errorf("invalid strategy range")
	}

	return start, end, nil
}
