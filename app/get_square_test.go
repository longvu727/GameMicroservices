package app

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	db "github.com/longvu727/FootballSquaresLibs/DB/db"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type GetGameTestSuite struct {
	suite.Suite
}

func (suite *GetGameTestSuite) SetupTest() {
}

func (suite *GetGameTestSuite) TestGetGame() {
	randomGame := randomGame()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetGame(gomock.Any(), gomock.Eq(randomGame.GameID)).
		Times(1).
		Return(randomGame, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	getGameParams := GetGameParams{GameID: randomGame.GameID}
	game, err := NewGameApp().GetDBGame(getGameParams, resources)
	suite.NoError(err)

	suite.Equal(randomGame.GameID, int32(game.GameID))
	suite.Equal(randomGame.GameGuid, game.GameGUID)
	suite.Equal(randomGame.Sport.String, game.Sport)
	suite.Equal(randomGame.TeamA.String, game.TeamA)
	suite.Equal(randomGame.TeamB.String, game.TeamB)
}

func (suite *GetGameTestSuite) TestGetGameByGUID() {
	randomGame := randomGameByGUID()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetGameByGUID(gomock.Any(), gomock.Eq(randomGame.GameGuid)).
		Times(1).
		Return(randomGame, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, context.Background())

	getGameParams := GetGameByGUIDParams{GameGUID: randomGame.GameGuid}
	game, err := NewGameApp().GetGameByGUID(getGameParams, resources)
	suite.NoError(err)

	suite.Equal(randomGame.GameID, int32(game.GameID))
	suite.Equal(randomGame.GameGuid, game.GameGUID)
	suite.Equal(randomGame.Sport.String, game.Sport)
	suite.Equal(randomGame.TeamA.String, game.TeamA)
	suite.Equal(randomGame.TeamB.String, game.TeamB)
}

func randomGame() db.GetGameRow {
	return db.GetGameRow{
		GameID:   rand.Int31n(1000),
		GameGuid: uuid.NewString(),
		Sport:    sql.NullString{String: "football", Valid: true},
		TeamA:    sql.NullString{String: "TeamA", Valid: true},
		TeamB:    sql.NullString{String: "TeamB", Valid: true},
	}
}

func randomGameByGUID() db.GetGameByGUIDRow {
	return db.GetGameByGUIDRow{
		GameID:   rand.Int31n(1000),
		GameGuid: uuid.NewString(),
		Sport:    sql.NullString{String: "football", Valid: true},
		TeamA:    sql.NullString{String: "TeamA", Valid: true},
		TeamB:    sql.NullString{String: "TeamB", Valid: true},
	}
}

func TestGetGameTestSuite(t *testing.T) {
	suite.Run(t, new(GetGameTestSuite))
}
