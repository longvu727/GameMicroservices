package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

type GetGameParams struct {
	GameID int32 `json:"game_id"`
}
type GetGameResponse struct {
	GameID   int32
	GameGuid string
	Sport    string
	TeamA    string
	TeamB    string

	ErrorMessage string `json:"error_message"`
}

func (response GetGameResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func GetDBGame(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*GetGameResponse, error) {
	var getGameResponse GetGameResponse
	var getGameParams GetGameParams
	json.NewDecoder(request.Body).Decode(&getGameParams)

	gameRow, err := dbConnect.QUERIES.GetGame(ctx, getGameParams.GameID)
	if err != nil {
		return &getGameResponse, err
	}

	getGameResponse.GameID = gameRow.GameID
	getGameResponse.GameGuid = gameRow.GameGuid
	getGameResponse.Sport = gameRow.Sport.String
	getGameResponse.TeamA = gameRow.TeamA.String
	getGameResponse.TeamB = gameRow.TeamB.String

	return &getGameResponse, nil
}
