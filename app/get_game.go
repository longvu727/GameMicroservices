package app

import (
	"encoding/json"

	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type GetGameParams struct {
	GameID int32 `json:"game_id"`
}

type GetGameByGUIDParams struct {
	GameGUID string `json:"game_guid"`
}

type GetGameResponse struct {
	services.Game
	ErrorMessage string `json:"error_message"`
}

func (response GetGameResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (game *GameApp) GetDBGame(getGameParams GetGameParams, resources *resources.Resources) (*GetGameResponse, error) {
	var getGameResponse GetGameResponse

	gameRow, err := resources.DB.GetGame(resources.Context, getGameParams.GameID)
	if err != nil {
		return &getGameResponse, err
	}

	getGameResponse.GameID = int64(gameRow.GameID)
	getGameResponse.GameGUID = gameRow.GameGuid
	getGameResponse.Sport = gameRow.Sport.String
	getGameResponse.TeamA = gameRow.TeamA.String
	getGameResponse.TeamB = gameRow.TeamB.String

	return &getGameResponse, nil
}

func (game *GameApp) GetGameByGUID(getGameByGUIDParams GetGameByGUIDParams, resources *resources.Resources) (*GetGameResponse, error) {
	var getGameResponse GetGameResponse

	gameRow, err := resources.DB.GetGameByGUID(resources.Context, getGameByGUIDParams.GameGUID)
	if err != nil {
		return &getGameResponse, err
	}

	getGameResponse.GameID = int64(gameRow.GameID)
	getGameResponse.GameGUID = gameRow.GameGuid
	getGameResponse.Sport = gameRow.Sport.String
	getGameResponse.TeamA = gameRow.TeamA.String
	getGameResponse.TeamB = gameRow.TeamB.String

	return &getGameResponse, nil
}
