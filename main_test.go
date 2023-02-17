package main

import (
	"testing"
)

func Test_rollDice(t *testing.T) {
	tests := []struct {
		name        string
		expectation func(int8) bool
	}{
		{
			name:        "Should return value between 1 and 6",
			expectation: func(result int8) bool { return result >= 1 && result <= 6 },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rollDice()
			if !tt.expectation(got) {
				t.Errorf("rollDice() = %v, which is not between 1 and 6", got)
			}
		})
	}
}

func Test_player_getDecision(t *testing.T) {
	tests := []struct {
		name string
		p    player
		want string
	}{
		{
			name: "Should return 'roll' when turn score is less than hold score",
			p: player{
				turnScore: 3,
				holdScore: 20,
			},
			want: "roll",
		},
		{
			name: "Should return 'hold' when turn score is greater than or equal to hold score",
			p: player{
				turnScore: 25,
				holdScore: 20,
			},
			want: "hold",
		},
		{
			name: "Should return 'hold' when player's total score plus turn score is greater than or equal to winning score",
			p: player{
				score:     97,
				turnScore: 6,
				holdScore: 20,
			},
			want: "hold",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.getDecision(); got != tt.want {
				t.Errorf("player.getDecision() = %v, want %v", got, tt.want)
			}
		})
	}
}
