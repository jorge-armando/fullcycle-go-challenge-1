package main

import (
	"fmt"
	"jorge/backend/routes"
	"jorge/env"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	address := fmt.Sprintf(":%s", env.SERVER_PORT)

	mux.HandleFunc("/", routes.CotacaoRoute)

	err := http.ListenAndServe(address, mux)

	if err != nil {
		panic(err)
	}
}
