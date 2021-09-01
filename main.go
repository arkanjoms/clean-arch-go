package main

import (
	"clean-arch-go/domain/service"
	"clean-arch-go/infra/database"
	"clean-arch-go/infra/factory"
	"clean-arch-go/infra/gateway/memory"
	"clean-arch-go/infra/http"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	db := database.GetInstance()
	repositoryFactory := factory.NewDatabaseRepositoryFactory(db)
	zipcodeClient := memory.NewZipcodeClient()
	freightCalculator := service.NewFreightCalculator()
	server := http.NewGorillaMux()
	http.NewRouteConfig(server, repositoryFactory, zipcodeClient, freightCalculator).Build()
	logrus.Fatal(server.Listen(8080))
}
