package tenniskata

type tennisGame3 struct {
	player1Score int
	player2Score int
	player1Name  string
	player2Name  string
}

func TennisGame3(player1Name string, player2Name string) TennisGame {
	game := &tennisGame3{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame3) GetScore() string {
	var s string
	if game.player1Score < 4 && game.player2Score < 4 && game.player1Score+game.player2Score != 6 {
		p := []string{"Love", "Fifteen", "Thirty", "Forty"}
		s = p[game.player1Score]
		if game.player1Score == game.player2Score {
			return s + "-All"
		}
		return s + "-" + p[game.player2Score]
	}
	if game.player1Score == game.player2Score {
		return "Deuce"
	}
	if game.player1Score > game.player2Score {
		s = game.player1Name
	} else {
		s = game.player2Name
	}
	if (game.player1Score-game.player2Score)*(game.player1Score-game.player2Score) == 1 {
		return "Advantage " + s
	}
	return "Win for " + s
}

func (game *tennisGame3) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.player1Score++
	} else {
		game.player2Score++
	}
}
