package main

import (
	"clean-arch-go/domain/service"
	"clean-arch-go/infra/database"
	"clean-arch-go/infra/factory"
	"clean-arch-go/infra/gateway/memory"
	"clean-arch-go/infra/http"
	"github.com/sirupsen/logrus"
    "github.com/subosito/gotenv"
    "os"
)

func main() {
    _ = gotenv.Load()
    level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
    if err != nil {
        logrus.Fatal("could not parse env LOG_LEVEL: %v", err)
    }
    logrus.SetLevel(level)
	db := database.GetInstance()
	repositoryFactory := factory.NewDatabaseRepositoryFactory(db)
	zipcodeClient := memory.NewZipcodeClient()
	freightCalculator := service.NewFreightCalculator()
	server := http.NewGorillaMux()
	http.NewRouteConfig(server, repositoryFactory, zipcodeClient, freightCalculator).Build()
	logrus.Fatal(server.Listen(8080))
}
