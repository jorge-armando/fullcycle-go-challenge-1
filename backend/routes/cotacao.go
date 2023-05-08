package routes

import (
	"encoding/json"
	"jorge/backend/repository"
	"jorge/backend/request"
	"net/http"
)

type CotacaoRouteResponse struct {
	CurrentQuote string `json:"current_quote"`
}

func CotacaoRoute(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json")

	lastQuote := request.GetLastQuote()

	response := CotacaoRouteResponse{
		CurrentQuote: lastQuote,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	repository.CotacaoRepo.Up()
	err = repository.CotacaoRepo.Create(lastQuote)

	if err != nil {
		panic(err)
	}

	w.Write(jsonResponse)
}
