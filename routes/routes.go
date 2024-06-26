package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/app"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

type Handler = func(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context)

func Register(db *db.MySQL, ctx context.Context) {
	log.Println("Registering routes")

	routes := map[string]Handler{
		"/":              home,
		"/CreateGame":    createGame,
		"/GetGame":       getGame,
		"/GetGameByGUID": getGameByGUID,
	}

	for route, handler := range routes {
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, db, ctx)
		})
	}
}

func home(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func createGame(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	createSquareResponse, err := app.CreateDBGame(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.GameGUID = ""
		createSquareResponse.ErrorMessage = `Unable to create game`
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func getGame(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	getGameResponse, err := app.GetDBGame(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getGameResponse.ErrorMessage = `Unable to get game`
		writer.Write(getGameResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getGameResponse.ToJson())
}

func getGameByGUID(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	getGameResponse, err := app.GetGameByGUID(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getGameResponse.ErrorMessage = `Unable to get game`
		writer.Write(getGameResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getGameResponse.ToJson())
}
