package main

import (
	"github.com/Kiveri/rating/cmd/service_provider"
	"github.com/Kiveri/rating/internal/config"
)

func main() {
	_ = service_provider.NewServiceProvider(config.NewConfig())
}
