package tenniskata

import "fmt"

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

func (game *tennisGame1) GetScore() string {

	array := [4]string{"Love", "Fifteen", "Thirty", "Forty"}

	if game.m_score1 == game.m_score2 {
		if game.m_score1 > 2 {
			return "Deuce"
		}
		return fmt.Sprintf("%s-All", array[game.m_score1])
	} else if game.m_score1 >= 4 || game.m_score2 >= 4 {
		minusResult := game.m_score1 - game.m_score2
		switch {
		case minusResult == 1:
			return fmt.Sprint("Advantage ", game.player1Name)
		case minusResult == -1:
			return fmt.Sprint("Advantage ", game.player2Name)
		case minusResult >= 2:
			return fmt.Sprint("Win for ", game.player1Name)
		default:
			return fmt.Sprint("Win for ", game.player2Name)
		}

	}

	return fmt.Sprintf("%s-%s", array[game.m_score1], array[game.m_score2])

}
