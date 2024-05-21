package main

import (
	"github.com/fajrulaulia/minder/config"
	"github.com/fajrulaulia/minder/src/delivery"
	"github.com/fajrulaulia/minder/src/repository"
	"github.com/fajrulaulia/minder/src/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	cfg := config.InitConfig()

	userRepos := repository.NewUserRepository(cfg)
	userUsecase := usecase.NewUserUsecase(userRepos)
	deliverUser := delivery.NewUserDelivery(userUsecase)

	matcherRepos := repository.NewMatcherRepository(cfg)
	matcherUsecase := usecase.NewMatcherUsecase(matcherRepos, userRepos)
	deliveryMatcher := delivery.NewMatcherDelivery(matcherUsecase)

	e := echo.New()

	deliverUser.Apply(e)
	deliveryMatcher.Apply(e)

	e.Logger.Fatal(e.Start(":8081"))

}
