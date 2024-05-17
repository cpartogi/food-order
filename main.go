package main

import (
	"fmt"
	appInit "food-order/init"

	handler "food-order/domain/food/handler"
	"food-order/domain/food/repo"
	"food-order/domain/food/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	appInit.StartAppInit()
}

func main() {
	db, err := appInit.ConnectToPGServerRead()

	if err != nil {
		fmt.Println("error connect to database")
	}

	defer db.DB.Close()

	//router
	e := echo.New()

	//middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	//depedency
	foodRepo := repo.NewFoodRepo(db.DB)

	authUc := usecase.NewFoodUsecase(foodRepo)

	handler.NewFoodHander(e, authUc)
	//e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
