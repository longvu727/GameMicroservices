package app

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type CreateGameParams struct {
	Sport      string `json:"sport"`
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

func (game *GameApp) CreateDBGame(createGameParams CreateGameParams, resources *resources.Resources) (*CreateGameResponse, error) {
	var createGameResponse CreateGameResponse

	gameGuid := (uuid.New()).String()

	gameID, err := resources.DB.CreateGame(resources.Context, db.CreateGameParams{
		GameGuid: gameGuid,
		Sport:    sql.NullString{String: "football", Valid: true},
		TeamA:    sql.NullString{String: createGameParams.TeamA, Valid: true},
		TeamB:    sql.NullString{String: createGameParams.TeamB, Valid: true},
	})

	if err != nil {
		return &createGameResponse, err
	}

	createGameResponse.GameGUID = gameGuid
	createGameResponse.GameID = gameID

	return &createGameResponse, err
}
