package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/app"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

type Handler = func(writer http.ResponseWriter, request *http.Request)

func Register(db *db.MySQL, ctx context.Context) {
	log.Println("Registering routes")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		home(w, r)
	})

	http.HandleFunc(http.MethodPost+" /CreateGame", func(w http.ResponseWriter, r *http.Request) {
		createGame(w, r, db, ctx)
	})

	http.HandleFunc(http.MethodPost+" /GetGame", func(w http.ResponseWriter, r *http.Request) {
		getGame(w, r, db, ctx)
	})
}

func home(writer http.ResponseWriter, _ *http.Request) {
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
