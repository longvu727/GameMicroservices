package routes

import (
	"bytes"
	"gamemicroservices/app"
	mockgameapp "gamemicroservices/app/mock"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type RoutesTestSuite struct {
	suite.Suite
}

func (suite *RoutesTestSuite) TestCreateGame() {

	url := "/CreateGame"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"sport":"football", "team_a": "red", "team_b": "blue"}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockGame := mockgameapp.NewMockGame(ctrl)
	mockGame.EXPECT().
		CreateDBGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateGameResponse{GameID: 10, GameGUID: uuid.NewString()}, nil)

	routes := Routes{Apps: mockGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetGame() {

	url := "/GetGame"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"game_id":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	returnGame := &app.GetGameResponse{}
	returnGame.GameID = 10
	returnGame.GameGUID = uuid.NewString()
	returnGame.Sport = "football"
	returnGame.TeamA = "red"
	returnGame.TeamB = "blue"

	mockGame := mockgameapp.NewMockGame(ctrl)
	mockGame.EXPECT().
		GetDBGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnGame, nil)

	routes := Routes{Apps: mockGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestHome() {

	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	routes := NewRoutes()
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
