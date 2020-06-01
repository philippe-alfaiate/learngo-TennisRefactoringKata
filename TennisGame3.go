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
	var s string
	if game.scorePlayer1 < 4 && game.scorePlayer2 < 4 && game.scorePlayer1+game.scorePlayer2 != 6 {
		p := []string{"Love", "Fifteen", "Thirty", "Forty"}
		s = p[game.scorePlayer1]
		if game.scorePlayer1 == game.scorePlayer2 {
			return s + "-All"
		}
		return s + "-" + p[game.scorePlayer2]
	}
	if game.scorePlayer1 == game.scorePlayer2 {
		return "Deuce"
	}
	if game.scorePlayer1 > game.scorePlayer2 {
		s = game.namePlayer1
	} else {
		s = game.namePlayer2
	}
	if (game.scorePlayer1-game.scorePlayer2)*(game.scorePlayer1-game.scorePlayer2) == 1 {
		return "Advantage " + s
	}
	return "Win for " + s
}

func (game *tennisGame3) WonPoint(namePlayer string) {
	if namePlayer == game.namePlayer1 {
		game.scorePlayer1++
	} else {
		game.scorePlayer2++
	}
}
