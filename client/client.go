package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"jorge/backend/routes"
	"jorge/env"
	"net/http"
)

func GetLastQuote() *routes.CotacaoRouteResponse {
	url := fmt.Sprintf("http://localhost:%s", env.SERVER_PORT)
	ctx, _ := context.WithTimeout(context.Background(), env.CLIENT_REQUEST_TIMEOUT)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var response routes.CotacaoRouteResponse

	err = json.Unmarshal(result, &response)

	if err != nil {
		panic(err)
	}

	return &response
}
