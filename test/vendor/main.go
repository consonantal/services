package main

import (
	// a relative import outside the go mod
	_ "github.com/micro/services/test/routes"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	srv := service.New(
		service.Name("vendor"),
	)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}