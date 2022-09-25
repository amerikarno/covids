package repository_test

import (
	"fmt"
	"testing"
	"wnexam/repository"
)

func TestGetJSONUrl(t *testing.T) {
	t.Parallel()
	t.Run("test get JSON", func(t *testing.T){
		url := "https://static.wongnai.com/devinterview/covid-cases.json"

		repo := repository.NewDataRepository(url)
		gju, err := repo.GetJSONUrl()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(gju)
	})

	t.Run("test url error", func(t *testing.T){
		url := "https://static.wongnai.com/devinterview/covid-cases.jso"

		repo := repository.NewDataRepository(url)
		gju, err := repo.GetJSONUrl()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(gju)
	})
}

func TestDecodeJSONUrl(t *testing.T){
	t.Parallel()
	t.Run("test decode JSON function", func(t *testing.T){
		url := "https://static.wongnai.com/devinterview/covid-cases.json"
		repo := repository.NewDataRepository(url)
		dju, err := repo.DecodeJSONUrl()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(dju)
	})
}