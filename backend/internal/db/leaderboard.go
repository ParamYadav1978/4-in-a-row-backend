package db

import "log"

func EnsurePlayer(username string) {

	// 🚫 NEVER insert BOT into leaderboard
	if username == "BOT" {
		return
	}

	query := `
	INSERT INTO leaderboard (username)
	VALUES ($1)
	ON CONFLICT (username) DO NOTHING;
	`
	_, err := DB.Exec(query, username)
	if err != nil {
		log.Println("EnsurePlayer error:", err)
	}
}


func RecordResult(winner string, loser string, isBot bool) {

	// Ensure only HUMAN players exist
	EnsurePlayer(winner)
	EnsurePlayer(loser)

	// ✅ Update winner (only if human)
	if winner != "BOT" {
		DB.Exec(`
		UPDATE leaderboard
		SET 
			wins = wins + 1,
			games_played = games_played + 1,
			updated_at = NOW()
		WHERE username = $1;
		`, winner)
	}

	// ✅ Update loser (only if human)
	if loser != "BOT" {
		DB.Exec(`
		UPDATE leaderboard
		SET 
			losses = losses + 1,
			games_played = games_played + 1,
			updated_at = NOW()
		WHERE username = $1;
		`, loser)
	}

	// 🤖 Bot game count → ONLY HUMAN
	if isBot {
		if winner != "BOT" {
			DB.Exec(`
			UPDATE leaderboard
			SET bot_games = bot_games + 1
			WHERE username = $1;
			`, winner)
		}

		if loser != "BOT" {
			DB.Exec(`
			UPDATE leaderboard
			SET bot_games = bot_games + 1
			WHERE username = $1;
			`, loser)
		}
	}
}



type LeaderboardRow struct {
	Username     string
	GamesPlayed  int
	Wins         int
	Losses       int
	Draws        int
	BotGames     int
}

func GetLeaderboard() ([]LeaderboardRow, error) {
	rows, err := DB.Query(`
		SELECT username, games_played, wins, losses, draws, bot_games
		FROM leaderboard
		ORDER BY wins DESC, games_played ASC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []LeaderboardRow

	for rows.Next() {
		var r LeaderboardRow
		err := rows.Scan(
			&r.Username,
			&r.GamesPlayed,
			&r.Wins,
			&r.Losses,
			&r.Draws,
			&r.BotGames,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}
