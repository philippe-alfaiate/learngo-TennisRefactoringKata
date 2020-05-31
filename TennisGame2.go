package tenniskata

import "math"

var scoreMap = [4]string{"Love", "Fifteen", "Thirty", "Forty"}

type tennisGame2 struct {
	P1point int
	P2point int

	P1res       string
	P2res       string
	player1Name string
	player2Name string
}

func TennisGame2(player1Name string, player2Name string) TennisGame {

	game := &tennisGame2{
		player1Name: player1Name,
		player2Name: player2Name,
		P1res:       scoreMap[0],
		P2res:       scoreMap[0],
	}

	return game
}

func (game *tennisGame2) GetEquality() string {

	if game.P1point < 3 {
		return game.P1res + "-All"
	}

	return "Deuce"
}

func (game *tennisGame2) GetWinner() string {

	if game.P1point-game.P2point > 0 {
		return "Win for " + game.player1Name
	}

	return "Win for " + game.player2Name
}

func (game *tennisGame2) GetAvantage() string {

	if game.P1point > game.P2point {
		return "Advantage " + game.player1Name
	}

	return "Advantage " + game.player2Name
}

func (game *tennisGame2) hasWinner() bool {
	return (game.P1point >= 4 || game.P2point >= 4) && math.Abs(float64(game.P1point-game.P2point)) >= 2
}
func (game *tennisGame2) hasAdvantage() bool {
	return game.P1point >= 3 && game.P2point >= 3 && game.P1point != game.P2point
}
func (game *tennisGame2) hasEquality() bool {
	return game.P1point == game.P2point
}

func (game *tennisGame2) GetScore() string {

	if game.hasWinner() {
		return game.GetWinner()
	}

	if game.hasAdvantage() {
		return game.GetAvantage()
	}

	if game.hasEquality() {
		return game.GetEquality()
	}

	return game.P1res + "-" + game.P2res
}

func (game *tennisGame2) UpdateP1Score() {

	game.P1point++

	if game.P1point < 4 {
		game.P1res = scoreMap[game.P1point]
	}
}

func (game *tennisGame2) UpdateP2Score() {

	game.P2point++

	if game.P2point < 4 {
		game.P2res = scoreMap[game.P2point]
	}
}

func (game *tennisGame2) WonPoint(player string) {
	if player == game.player1Name {
		game.UpdateP1Score()
	} else {
		game.UpdateP2Score()
	}
}
