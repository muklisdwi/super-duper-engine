//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire

package main

import (
	"fmt"
	"go-boiler/config"
	"go-boiler/utils/log"
)

var (
	configuration *config.Config
)

func main() {
	// initialize log
	log.InitializeLog()

	// initialize configuration
	configuration = config.InitializeConfiguration()

	// setup log
	log.SetLogLevel(configuration)

	// initialize service
	service := InitializeService()

	// start service
	service.SetupAndServe()

	fmt.Println("First Init :D")
}
