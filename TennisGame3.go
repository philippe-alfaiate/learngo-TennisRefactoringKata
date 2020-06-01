package tenniskata

type tennisGame3 struct {
	scorePlayer1 int
	scorePlayer2 int
	namePlayer1  string
	namePlayer2  string
}

func TennisGame3(namePlayer1 string, namePlayer2 string) TennisGame {
	game := &tennisGame3{
		namePlayer1: namePlayer1,
		namePlayer2: namePlayer2}

	return game
}

func (game *tennisGame3) GetScore() string {
	var score string
	if game.scorePlayer1 < 4 && game.scorePlayer2 < 4 && game.scorePlayer1+game.scorePlayer2 != 6 {
		scoreMap := []string{"Love", "Fifteen", "Thirty", "Forty"}
		score = scoreMap[game.scorePlayer1]
		if game.scorePlayer1 == game.scorePlayer2 {
			return score + "-All"
		}
		return score + "-" + scoreMap[game.scorePlayer2]
	}
	if game.scorePlayer1 == game.scorePlayer2 {
		return "Deuce"
	}
	if game.scorePlayer1 > game.scorePlayer2 {
		score = game.namePlayer1
	} else {
		score = game.namePlayer2
	}
	if game.scorePlayer1-game.scorePlayer2 == 1 || game.scorePlayer2-game.scorePlayer1 == 1 {
		return "Advantage " + score
	}
	return "Win for " + score
}

func (game *tennisGame3) WonPoint(namePlayer string) {
	if namePlayer == game.namePlayer1 {
		game.scorePlayer1++
	} else {
		game.scorePlayer2++
	}
}
