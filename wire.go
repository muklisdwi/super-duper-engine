//go:build wireinject
// +build wireinject

package main

import (
	"go-boiler/config"
	"go-boiler/server/http"

	"github.com/google/wire"
)

func InitializeService() *http.ServerHTTP {
	wire.Build(
		config.InitializeConfiguration,
		http.InitializeServer,
	)
	return nil
}
