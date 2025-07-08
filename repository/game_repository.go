package repository

import (
	"database/sql"
	"errors"

	"github.com/feribeirods/api-go-gin/model"
)

// Insert the values recieved in the DB
func InsertGame(db *sql.DB, game *model.Game) error {
	query := `
		INSERT INTO games (id, name, publish_date, beated, beated_at, rating)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(
		query,
		game.ID,
		game.Name,
		game.PublishDate,
		game.Beated,
		game.BeatedAt,
		game.Rating,
	)

	return err
}

// Get all games stores on the DB
func FindAllGames(db *sql.DB) ([]model.Game, error) {
	rows, err := db.Query(`SELECT id, name, publish_date, beated, beated_at, rating FROM games`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []model.Game

	for rows.Next() {
		var g model.Game
		err := rows.Scan(&g.ID, &g.Name, &g.PublishDate, &g.Beated, &g.BeatedAt, &g.Rating)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

// Delete an especific game by the ID
func DeleteGameByID(db *sql.DB, id int64) error {
	result, err := db.Exec(`DELETE FROM games WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}

// Update the game passed by ID
func UpdateGameByID(db *sql.DB, id int64, game *model.Game) error {
	query := `
		UPDATE games
		SET name = ?, publish_date = ?, beated = ?, beated_at = ?, rating = ?
		WHERE id = ?
	`

	result, err := db.Exec(
		query,
		game.Name,
		game.PublishDate,
		game.Beated,
		game.BeatedAt,
		game.Rating,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}
