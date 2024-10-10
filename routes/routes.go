package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gamemicroservices/app"
	"log"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type RoutesInterface interface {
	Register(resources *resources.Resources) *http.ServeMux
}

type Routes struct {
	Apps app.Game
}

type Handler = func(writer http.ResponseWriter, request *http.Request, resources *resources.Resources)

func NewRoutes() RoutesInterface {
	return &Routes{
		Apps: app.NewGameApp(),
	}
}
func (routes *Routes) Register(resources *resources.Resources) *http.ServeMux {
	log.Println("Registering routes")
	mux := http.NewServeMux()

	routesHandlersMap := map[string]Handler{
		"/":                                 routes.home,
		http.MethodPost + " /CreateGame":    routes.createGame,
		http.MethodPost + " /GetGame":       routes.getGame,
		http.MethodPost + " /GetGameByGUID": routes.getGameByGUID,
	}

	for route, handler := range routesHandlersMap {
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, resources)
		})
	}

	return mux
}

func (routes *Routes) home(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func (routes *Routes) createGame(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var createGameParams app.CreateGameParams
	json.NewDecoder(request.Body).Decode(&createGameParams)

	createSquareResponse, err := routes.Apps.CreateDBGame(createGameParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.GameGUID = ""
		createSquareResponse.ErrorMessage = `Unable to create game` + err.Error()
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func (routes *Routes) getGame(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getGameParams app.GetGameParams
	json.NewDecoder(request.Body).Decode(&getGameParams)

	getGameResponse, err := routes.Apps.GetDBGame(getGameParams, resources)

	if err != nil && err == sql.ErrNoRows {
		getGameResponse.ErrorMessage = `Game not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getGameResponse.ErrorMessage = `Unable to get game` + err.Error()
		writer.Write(getGameResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getGameResponse.ToJson())
}

func (routes *Routes) getGameByGUID(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getGameByGUIDParams app.GetGameByGUIDParams
	json.NewDecoder(request.Body).Decode(&getGameByGUIDParams)

	log.Printf("Received request for %s\n", getGameByGUIDParams)

	getGameResponse, err := routes.Apps.GetGameByGUID(getGameByGUIDParams, resources)

	if err != nil && err == sql.ErrNoRows {
		getGameResponse.ErrorMessage = `Game not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getGameResponse.ErrorMessage = `Unable to get game` + err.Error()
		writer.Write(getGameResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getGameResponse.ToJson())
}
