package app

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type CreateGameTestSuite struct {
	suite.Suite
}

func TestCreateGameTestSuite(t *testing.T) {
	suite.Run(t, new(CreateGameTestSuite))
}

func (suite *CreateGameTestSuite) TestCreateGame() {
	randomGame := randomGame()

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		CreateGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(int64(randomGame.GameID), nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	createSquareParams := CreateGameParams{
		Sport: randomGame.Sport.String,
		TeamA: randomGame.TeamA.String,
		TeamB: randomGame.TeamB.String,
	}
	game, err := NewGameApp().CreateDBGame(createSquareParams, resources)
	suite.NoError(err)

	suite.Equal(randomGame.GameID, int32(game.GameID))
	suite.Greater(len(game.ToJson()), 0)
}

func (suite *CreateGameTestSuite) TestCreateGameDBError() {

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		CreateGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(int64(0), errors.New("test error"))

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	_, err = NewGameApp().CreateDBGame(CreateGameParams{}, resources)
	suite.Error(err)
}
