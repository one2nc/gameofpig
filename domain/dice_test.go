package domain

import "testing"

func TestDice_Roll(t *testing.T) {
	tests := []struct {
		name        string
		d           Dice
		expectation func(int8) bool
	}{
		{
			name:        "Should return value between 1 and 6",
			expectation: func(result int8) bool { return result >= 1 && result <= 6 },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dice{}
			got := d.Roll()
			if !tt.expectation(got) {
				t.Errorf("rollDice() = %v, which is not between 1 and 6", got)
			}
		})
	}
}
