package tenniskata

import (
	"fmt"
)

type tennisGame1 struct {
	m_score1    int
	m_score2    int
	player1Name string
	player2Name string
}

//TennisGame1 is the first test
func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.m_score1++
	} else {
		game.m_score2++
	}
}

//Constante
const (
	advantage  string = "Advantage"
	win               = "Win for"
	all               = "All"
	score0            = "Love"
	score1            = "Fifteen"
	score2            = "Thirty"
	score3            = "Forty"
	scoreDeuce        = "Deuce"
)

//GetFormatedScore Format output string of score
func GetFormatedScore(fisrt string, second string, separator ...string) string {
	def := " "
	if separator != nil {
		def = separator[0]
	}
	return fmt.Sprint(fisrt, def, second)
}

func (game *tennisGame1) GetScore() string {

	score := [4]string{score0, score1, score2, score3}

	if game.m_score1 == game.m_score2 {
		if game.m_score1 > 2 {
			return scoreDeuce
		}
		return GetFormatedScore(score[game.m_score1], all, "-")
	}
	if game.m_score1 >= 4 || game.m_score2 >= 4 {
		minusResult := game.m_score1 - game.m_score2
		switch {
		case minusResult == 1:
			return GetFormatedScore(advantage, game.player1Name)
		case minusResult == -1:
			return GetFormatedScore(advantage, game.player2Name)
		case minusResult >= 2:
			return GetFormatedScore(win, game.player1Name)
		default:
			return GetFormatedScore(win, game.player2Name)
		}
	}

	return GetFormatedScore(score[game.m_score1], score[game.m_score2], "-")
}
