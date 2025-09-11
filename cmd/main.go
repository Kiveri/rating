package main

import (
	"net/http"

	"github.com/Kiveri/rating/cmd/service_provider"
	"github.com/Kiveri/rating/internal/config"
)

//go run cmd/main.go -storage=postgres

func main() {
	sp := service_provider.NewServiceProvider(config.NewConfig())

	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}
