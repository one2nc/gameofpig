package domain

type Player struct {
	Name      string
	HoldScore int8
	Score     int8
	TurnScore int8
}

func (p Player) GetDecision(winningScore int8) string {
	if p.Score+p.TurnScore >= winningScore {
		return "hold"
	} else if p.TurnScore >= p.HoldScore {
		return "hold"
	} else {
		return "roll"
	}
}

func (p *Player) PlayTurn(dice Dice, winningScore int8) int8 {
	p.TurnScore = 0
	for p.GetDecision(winningScore) == "roll" {
		diceValue := dice.Roll()
		if diceValue == 1 {
			return 0
		}
		p.TurnScore += diceValue
	}
	return p.TurnScore
}
