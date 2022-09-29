package main

import (
	"github.com/amerikarno/covids/app"
	"github.com/amerikarno/covids/config"
	"github.com/amerikarno/covids/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Init()
}

func main() {
	handler := initHandler(config.App.URL)

	//start gin and route
	r := gin.Default()

	r.GET("/covid/summary", handler.GetCovidSum)

	r.Run()
}

func initHandler(url string) *app.Handler {
	repo := repository.NewDataRepository(url)
	apps := app.NewUsecase(repo)

	return app.NewHandler(*apps)
}
