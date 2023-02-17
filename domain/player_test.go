package domain

import "testing"

var winningScore = int8(100)

func TestPlayer_GetDecision(t *testing.T) {
	tests := []struct {
		name string
		p    Player
		want string
	}{
		{
			name: "Should return 'roll' when turn score is less than hold score",
			p: Player{
				TurnScore: 3,
				HoldScore: 20,
			},
			want: "roll",
		},
		{
			name: "Should return 'hold' when turn score is greater than or equal to hold score",
			p: Player{
				TurnScore: 25,
				HoldScore: 20,
			},
			want: "hold",
		},
		{
			name: "Should return 'hold' when player's total score plus turn score is greater than or equal to winning score",
			p: Player{
				Score:     97,
				TurnScore: 6,
				HoldScore: 20,
			},
			want: "hold",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetDecision(winningScore); got != tt.want {
				t.Errorf("player.getDecision() = %v, want %v", got, tt.want)
			}
		})
	}
}
