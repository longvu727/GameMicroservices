// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: games.sql

package db

import (
	"context"
	"database/sql"
)

const createGames = `-- name: CreateGames :execresult
INSERT INTO games (
  game_guid, sport, team_a, team_b
) VALUES (
  ?, ?, ?, ?
)
`

type CreateGamesParams struct {
	GameGuid string
	Sport    sql.NullString
	TeamA    sql.NullString
	TeamB    sql.NullString
}

func (q *Queries) CreateGames(ctx context.Context, arg CreateGamesParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createGames,
		arg.GameGuid,
		arg.Sport,
		arg.TeamA,
		arg.TeamB,
	)
}
