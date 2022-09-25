package main

import (
	"wnexam/app"
	"wnexam/repository"

	"github.com/gin-gonic/gin"
	// "github.com/rs/cors/wrapper/gin"
)

func main(){
	//url to get covid cases
	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	handler := initHandler(url)
	
	//start gin and route
	r := gin.Default()

	r.GET("/covid/summary", handler.GetCovidSum)
	
	r.Run()
}

func initHandler(url string) *app.Handler{
	repo := repository.NewDataRepository(url)
	apps := app.NewUsecase(repo)

	return app.NewHandler(*apps)
}
