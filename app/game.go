package app

import "github.com/longvu727/FootballSquaresLibs/util/resources"

type Game interface {
	GetDBGame(getGameParams GetGameParams, resources *resources.Resources) (*GetGameResponse, error)
	GetGameByGUID(getGameByGUIDParams GetGameByGUIDParams, resources *resources.Resources) (*GetGameResponse, error)
	CreateDBGame(createGameParams CreateGameParams, resources *resources.Resources) (*CreateGameResponse, error)
}

type GameApp struct{}

func NewGameApp() Game {
	return &GameApp{}
}
