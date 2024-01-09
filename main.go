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

	fmt.Println("First Init :D")
}
