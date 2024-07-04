// Code generated by MockGen. DO NOT EDIT.
// Source: gamemicroservices/app (interfaces: Game)

// Package mockgameapp is a generated GoMock package.
package mockgameapp

import (
	app "gamemicroservices/app"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/longvu727/FootballSquaresLibs/util/resources"
)

// MockGame is a mock of Game interface.
type MockGame struct {
	ctrl     *gomock.Controller
	recorder *MockGameMockRecorder
}

// MockGameMockRecorder is the mock recorder for MockGame.
type MockGameMockRecorder struct {
	mock *MockGame
}

// NewMockGame creates a new mock instance.
func NewMockGame(ctrl *gomock.Controller) *MockGame {
	mock := &MockGame{ctrl: ctrl}
	mock.recorder = &MockGameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGame) EXPECT() *MockGameMockRecorder {
	return m.recorder
}

// CreateDBGame mocks base method.
func (m *MockGame) CreateDBGame(arg0 app.CreateGameParams, arg1 *resources.Resources) (*app.CreateGameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDBGame", arg0, arg1)
	ret0, _ := ret[0].(*app.CreateGameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDBGame indicates an expected call of CreateDBGame.
func (mr *MockGameMockRecorder) CreateDBGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDBGame", reflect.TypeOf((*MockGame)(nil).CreateDBGame), arg0, arg1)
}

// GetDBGame mocks base method.
func (m *MockGame) GetDBGame(arg0 app.GetGameParams, arg1 *resources.Resources) (*app.GetGameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBGame", arg0, arg1)
	ret0, _ := ret[0].(*app.GetGameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBGame indicates an expected call of GetDBGame.
func (mr *MockGameMockRecorder) GetDBGame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBGame", reflect.TypeOf((*MockGame)(nil).GetDBGame), arg0, arg1)
}

// GetGameByGUID mocks base method.
func (m *MockGame) GetGameByGUID(arg0 app.GetGameByGUIDParams, arg1 *resources.Resources) (*app.GetGameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameByGUID", arg0, arg1)
	ret0, _ := ret[0].(*app.GetGameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameByGUID indicates an expected call of GetGameByGUID.
func (mr *MockGameMockRecorder) GetGameByGUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameByGUID", reflect.TypeOf((*MockGame)(nil).GetGameByGUID), arg0, arg1)
}
