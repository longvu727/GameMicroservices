package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

type CreateGameParams struct {
	Sport      string `json:"sport"`
	SquareSize int32  `json:"square_size"`
	TeamA      string `json:"team_a"`
	TeamB      string `json:"team_b"`
}
type CreateGameResponse struct {
	GameGUID     string `json:"game_guid"`
	GameID       int64  `json:"game_id"`
	ErrorMessage string `json:"error_message"`
}

func (response CreateGameResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func CreateDBGame(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*CreateGameResponse, error) {
	var createSquareParams CreateGameParams
	json.NewDecoder(request.Body).Decode(&createSquareParams)

	var createGameResponse CreateGameResponse

	gameGuid := (uuid.New()).String()

	gameID, err := insertGame(ctx, dbConnect, gameGuid, createSquareParams.TeamA, createSquareParams.TeamB)
	if err != nil {
		return &createGameResponse, err
	}

	createGameResponse.GameGUID = gameGuid
	createGameResponse.GameID = gameID

	return &createGameResponse, err
}

func insertGame(ctx context.Context, dbConnect *db.MySQL, gameGuid string, teamA string, teamB string) (int64, error) {

	createGameResult, err := dbConnect.QUERIES.CreateGames(ctx, db.CreateGamesParams{
		GameGuid: gameGuid,
		Sport:    sql.NullString{String: "football", Valid: true},
		TeamA:    sql.NullString{String: teamA, Valid: true},
		TeamB:    sql.NullString{String: teamB, Valid: true},
	})
	if err != nil {
		return 0, err
	}

	gameID, err := createGameResult.LastInsertId()
	if err != nil {
		return 0, err
	}
	return gameID, nil
}
