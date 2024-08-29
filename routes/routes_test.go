package routes

import (
	"bytes"
	"errors"
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

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}

func (suite *RoutesTestSuite) getTestError() error {
	return errors.New("test error")
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

func (suite *RoutesTestSuite) TestCreateGameError() {

	url := "/CreateGame"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockGame := mockgameapp.NewMockGame(ctrl)
	mockGame.EXPECT().
		CreateDBGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateGameResponse{}, suite.getTestError())

	routes := Routes{Apps: mockGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
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

func (suite *RoutesTestSuite) TestGetGameError() {

	url := "/GetGame"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockGame := mockgameapp.NewMockGame(ctrl)
	mockGame.EXPECT().
		GetDBGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.GetGameResponse{}, suite.getTestError())

	routes := Routes{Apps: mockGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
}

func (suite *RoutesTestSuite) TestGetGameByGUID() {
	testGuid := "f838b751-2553-46bc-a19a-cfb3bbac49a5"

	url := "/GetGameByGUID"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"game_guid":`+testGuid+`}`)))
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
		GetGameByGUID(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnGame, nil)

	routes := Routes{Apps: mockGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetGameByGUIDError() {

	url := "/GetGameByGUID"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockGame := mockgameapp.NewMockGame(ctrl)
	mockGame.EXPECT().
		GetGameByGUID(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.GetGameResponse{}, suite.getTestError())

	routes := Routes{Apps: mockGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusInternalServerError)
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
